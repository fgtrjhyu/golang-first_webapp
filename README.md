# Go言語でウェブアプリケーションを作るための練習
- githubに合わせてmasterではなくmainにする
```bash
git config --global init.defaultBranch main
```
- githubで作成したリポジトリをremoteのoriginに設定する
```bash
git remote add oigin git&commat;github.com:*username*/*repositoryname*.git
```
- markdownファイルはglowで読みます
```bash
sudo apt install -y glow
glow README.md
```
- mkcertで作った鍵(*.pem)を入れてはいけません
```bash
echo '*.pem' >> .gitignore
```
- Goのバージョン更新
```bash
goenv install 1.25.3
goenv local 1.25.3
go version
```
- Goの現在の依存関係を解消する
```bash
go mod tidy
```
- Goの依存関係を更新するので依存関係も変化する可能性がある
```bash
go get -u ./...
```
- 前のgo mod tidyでの更新を更に反映する
```bash
go mod tidy
```
- ビルドとテストを実行 
```bash
go build ./...
go test ./...
```
