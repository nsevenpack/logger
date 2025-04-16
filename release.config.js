module.exports = {
  branches: ['main'],
  plugins: [
    '@semantic-release/commit-analyzer',
    '@semantic-release/release-notes-generator',
    '@semantic-release/changelog',
    ['@semantic-release/github', {
      successComment: false,
      failComment: false,
      labels: [],
      releasedLabels: []
    }],
    '@semantic-release/git'
  ]
}
