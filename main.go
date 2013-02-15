package main

import (
        "fmt"
        // "flag"
        "encoding/json"
        "io"
        "net/http"
        "os"
        "flag"
        "regexp"
)

type Track struct {
        Track, Artist, Title string
}

func parseTracks(input io.Reader) []Track {
        fmt.Sprintf("%s", input)
        decoder := json.NewDecoder(input)

        var xs []Track
        decoder.Decode(&xs)
        return xs
}

func channelIdFor(genre string) int {
        var id int

        switch genre {
        case "psy":
                id = 8
        case "goa":
                id = 8
        case "goapsy":
                id = 8
        case "psytrance":
                id = 8
        case "electrohouse":
                id = 56
        case "eh":
                id = 56
        case "liquiddnb":
                id = 105
        case "ldnb":
                id = 105
        case "epictrance":
                id = 175
        case "et":
                id = 175
        case "handsup":
                id = 176
        case "hu":
                id = 176
        case "clubdubstep":
                id = 177
        case "clubds":
                id = 177
        case "glitchhop":
                id = 198
        case "gh":
                id = 198
        default:
                id = 175
        }

        return id
}

func channelNameFor(genre string) string {
        var result string

        switch genre {
        case "psy":
                result = "Goa/Psy Trance"
        case "goa":
                result = "Goa/Psy Trance"
        case "goapsy":
                result = "Goa/Psy Trance"
        case "psytrance":
                result = "Goa/Psy Trance"
        case "liquiddnb":
                result = "Liquid DnB"
        case "ldnb":
                result = "Liquid DnB"
        case "epictrance":
                result = "Epic Trance"
        case "et":
                result = "Epic Trance"
        case "electrohouse":
                result = "Electro House"
        case "eh":
                result = "Electro House"
        case "handsup":
                result = "Hands Up"
        case "hu":
                result = "Hands Up"
        case "clubdubstep":
                result = "Club Dubstep"
        case "clubds":
                result = "Club Dubstep"
        case "glitchhop":
                result = "Glitch Hop"
        case "gh":
                result = "Glitch Hop"
        default:
                result = "Epic Trance"
        }

        return result
}

func recentTracksURLFor(genre string) string {
        return fmt.Sprintf("http://api.audioaddict.com/v1/di/track_history/channel/%d.json", channelIdFor(genre))
}

func recentTracksFor(genre string) ([]Track, error) {
        var url = recentTracksURLFor(genre)

        fmt.Printf("GET %s...\n", url)
        resp, err := http.Get(url)

        var tracks []Track

        if err != nil {
                fmt.Printf("GET %s failed: %s", url, err)
        } else {
                defer resp.Body.Close()

                tracks = parseTracks(resp.Body)

                if err != nil {
                        return nil, err
                }
        }

        return tracks, nil
}

var adTrackTitle, _ = regexp.Compile(`^TSTAG.*`)
var separatorRe, _ = regexp.Compile(`\s-\s`)

func displayAdTrack(t Track) {
        fmt.Print("* /advertisement/\n")
}

func displayMusicTrack(t Track) {
        a := t.Artist
        ti := t.Title
        tr := separatorRe.ReplaceAllString(t.Track, " — ")

        if len(a) > 0 {
                fmt.Printf("* %s — %s\n", a, ti)
        } else {
                fmt.Printf("* %s\n", tr)
        }       
}

func displayTrack(t Track) {
        tr := t.Track

        if adTrackTitle.MatchString(tr) == true {
                displayAdTrack(t)
        } else {
                displayMusicTrack(t)
        }
}

func displayTracks(xs []Track, limit int, channel string) {
        fmt.Printf("\nRecent %s tracks:\n\n", channelNameFor(channel))

        var l int
        if len(xs) > limit {
                l = limit
        } else {
                l = len(xs)
        }

        for i := 0; i < l; i++ {
                displayTrack(xs[i])
        }
        fmt.Printf("\n")
}

var channel = flag.String("channel", "epictrance", "di.fm channel to use")
var limit   = flag.Int("limit", 5, "How many recent tracks to display")

func main() {
        flag.Parse()
        var tracks, err = recentTracksFor(*channel)

        if err != nil {
                fmt.Printf("Got an error: %s", err)
                os.Exit(1)
        } else {
                displayTracks(tracks, *limit, *channel)
        }

        os.Exit(0)
}
