module.exports = {
    branches: ['main'],
    plugins: [
      '@semantic-release/commit-analyzer', // détecte le type de version
      '@semantic-release/release-notes-generator',
      '@semantic-release/changelog',
      '@semantic-release/github',
      '@semantic-release/git'
    ]
  }
  