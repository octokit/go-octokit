package octokit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleasesService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/jingweno/gh/releases", "releases", nil)

	url, err := ReleasesURL.Expand(M{"owner": "jingweno", "repo": "gh"})
	assert.NoError(t, err)

	releases, result := client.Releases(url).All()
	assert.False(t, result.HasError())
	assert.Len(t, releases, 1)

	firstRelease := releases[0]
	assert.EqualValues(t, 50013, firstRelease.ID)
	assert.EqualValues(t, "v0.23.0", firstRelease.TagName)
	assert.EqualValues(t, "master", firstRelease.TargetCommitish)
	assert.EqualValues(t, "v0.23.0", firstRelease.Name)
	assert.False(t, firstRelease.Draft)
	assert.False(t, firstRelease.Prerelease)
	assert.EqualValues(t, "* Windows works!: https://github.com/jingweno/gh/commit/6cb80cb09fd9f624a64d85438157955751a9ac70", firstRelease.Body)
	assert.EqualValues(t, "https://api.github.com/repos/jingweno/gh/releases/50013", firstRelease.URL)
	assert.EqualValues(t, "https://api.github.com/repos/jingweno/gh/releases/50013/assets", firstRelease.AssetsURL)
	assert.EqualValues(t, "https://uploads.github.com/repos/jingweno/gh/releases/50013/assets{?name}", string(firstRelease.UploadURL))
	assert.EqualValues(t, "https://github.com/jingweno/gh/releases/v0.23.0", firstRelease.HTMLURL)
	assert.EqualValues(t, "2013-09-23 00:59:10 +0000 UTC", firstRelease.CreatedAt.String())
	assert.EqualValues(t, "2013-09-23 01:07:56 +0000 UTC", firstRelease.PublishedAt.String())

	firstReleaseAssets := firstRelease.Assets
	assert.Len(t, firstReleaseAssets, 8)

	firstAsset := firstReleaseAssets[0]
	assert.EqualValues(t, 20428, firstAsset.ID)
	assert.EqualValues(t, "gh_0.23.0-snapshot_amd64.deb", firstAsset.Name)
	assert.EqualValues(t, "gh_0.23.0-snapshot_amd64.deb", firstAsset.Label)
	assert.EqualValues(t, "application/x-deb", firstAsset.ContentType)
	assert.EqualValues(t, "uploaded", firstAsset.State)
	assert.EqualValues(t, 1562984, firstAsset.Size)
	assert.EqualValues(t, 0, firstAsset.DownloadCount)
	assert.EqualValues(t, "https://api.github.com/repos/jingweno/gh/releases/assets/20428", firstAsset.URL)
	assert.EqualValues(t, "2013-09-23 01:05:20 +0000 UTC", firstAsset.CreatedAt.String())
	assert.EqualValues(t, "2013-09-23 01:07:56 +0000 UTC", firstAsset.UpdatedAt.String())
}

func TestCreateRelease(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/Hello-World/releases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, "{\"tag_name\":\"v1.0.0\",\"target_commitish\":\"master\"}\n")
		respondWithJSON(w, loadFixture("create_release.json"))
	})

	url, err := ReleasesURL.Expand(M{"owner": "octokit", "repo": "Hello-World"})
	assert.NoError(t, err)

	params := Release{
		TagName:         "v1.0.0",
		TargetCommitish: "master",
	}
	release, result := client.Releases(url).Create(params)

	assert.False(t, result.HasError())
	assert.EqualValues(t, "v1.0.0", release.TagName)
}

func TestUpdateRelease(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/Hello-World/releases/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testBody(t, r, "{\"tag_name\":\"v1.0.0\",\"target_commitish\":\"master\"}\n")
		respondWithJSON(w, loadFixture("create_release.json"))
	})

	url, err := ReleasesURL.Expand(M{"owner": "octokit", "repo": "Hello-World", "id": "123"})
	assert.NoError(t, err)

	params := Release{
		TagName:         "v1.0.0",
		TargetCommitish: "master",
	}
	release, result := client.Releases(url).Update(params)

	assert.False(t, result.HasError())
	assert.EqualValues(t, "v1.0.0", release.TagName)
}
