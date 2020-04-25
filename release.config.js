module.exports = {
  branches: "master",
  repositoryUrl: "https://github.com/karl-cardenas-coding/disaster-cli.git",
  plugins: ["@semantic-release/commit-analyzer",
   "@semantic-release/release-notes-generator",
   ["@semantic-release/github", {
      "assets": ["disaster-**"]
      }
  ],
]
}
