/*
    Copyright (C) 2025  Oleksandr Miliukhin

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"github.com/go-rod/rod"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: noise-pollution <word> <client-id> <client-secret> \nUse \"\" as <word> to match every playlist.")
		return
	}
	word := os.Args[1]
	client_id := os.Args[2]
	client_secret := os.Args[3]
	fmt.Println(client_id)
	fmt.Println(client_secret)

	browser := rod.New().MustConnect()
	defer browser.MustClose()
	page := browser.MustPage("https://everynoise.com/everynoise1d-name.html")
	page.MustWaitLoad()

	file, err := os.OpenFile("links.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var playlistURLs []string

	trs := page.MustElements("tr")

	for _, tr := range trs {
		tds := tr.MustElements("td")

		if len(tds) < 3 {
			continue
		}

		thirdTdText := tds[2].MustText()

		if strings.Contains(thirdTdText, word) {
			playlistLink, err := tds[1].MustElement("a").Attribute("href")
			if err != nil {
				fmt.Println("Error retrieving the href:", err)
				continue
			}

			spotifyLink := *playlistLink
			parts := strings.Split(spotifyLink, "spotify:playlist:")
			if len(parts) > 1 {
				playlistID := parts[1]

				newURL := fmt.Sprintf("https://open.spotify.com/playlist/%s", playlistID)

				_, err := file.WriteString(newURL + " " + thirdTdText + "\n")
				if err != nil {
					fmt.Println("Error writing to file:", err)
					continue
				}

				playlistURLs = append(playlistURLs, newURL)
				fmt.Println("Playlist Link:", newURL)
			}
		}
	}
	fmt.Println("Saved all the links to links.txt. Downloading music now!")
	for _, url := range playlistURLs {
		cmd := exec.Command("bash", "-c", fmt.Sprintf("export SPOTIPY_CLIENT_ID='%s' && export SPOTIPY_CLIENT_SECRET='%s' && spotify_dl -w -l '%s'", client_id, client_secret, url))
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error downloading playlist %s: %v", url, err)
		}
	}
}
