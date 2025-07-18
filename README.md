# get-last-commit

Action for get last commit from GitHub repository

Usage example:
```yaml
name: print commit

jobs:
  build:
    name: Build binaries
    runs-on: ubuntu-latest
    steps:
      - name: Get commit
        uses: ci-space/get-last-commit@v0.1.0
        id: commit
        env:
          GITHUB_TOKEN: ${{ secrets.DEP_TOKEN }}

      - name: Print commit
        run: echo ${{ steps.commit.outputs.sha }}
```
