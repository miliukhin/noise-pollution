# noise-pollution - download playlists of everynoise.com

Searches everynoise.com for playlists containing the word provided as an
argument and downloads them. Can be used to download all the playlists at once.
Requires [spotify_dl](https://github.com/SathyaBhat/spotify-dl).

Playlists are downloaded one at a time because Spotify limits the number of
requests to its API. All the playlist links are saved to links.txt before
the downloads start, so you can try and use multiple Spotify apps or other
methods to download music faster.

## Usage

Create a Spotify app at https://developer.spotify.com/dashboard
(for detailed instructions check out spotify_dl docs) and get app's client id and secret.

```
usage: noise-pollution <word> <client-id> <client-secret>
Use "" as <word> to match every playlist.
```
