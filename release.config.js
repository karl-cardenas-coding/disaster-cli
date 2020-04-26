module.exports = {
  branches: ["master", "fix-1"],
  repositoryUrl: "https://github.com/karl-cardenas-coding/disaster-cli.git",
  plugins: ["@semantic-release/commit-analyzer",
   "@semantic-release/release-notes-generator",
   ["@semantic-release/exec", {
    "analyzeCommitsCmd": "echo 'NEW_VERSION=false' > VERSION.env",
    "verifyReleaseCmd": "echo 'export VERSION=${nextRelease.version}\nNEW_VERSION=true' > VERSION.env"
  }],
   ["@semantic-release/github", {
      "assets": ["*.zip"]
      }
  ],
]
}
