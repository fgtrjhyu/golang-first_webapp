# Go言語でウェブアプリケーションを作るための練習
- githubに合わせてmasterではなくmainにする
```bash
git config --global init.defaultBranch main
```
- githubで作成したリポジトリをremoteのoriginに設定する
```bash
git remote add oigin git&commat;github.com:*username*/*repositoryname*.git
```
- mkcertでTLSの鍵を作成する方法
```bash
mkcert \
  -cert-file certs/localhost+1.pem \
  -key-file certs/localhost+1-key.pem \
  localhost 127.0.0.1 ::1
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
# 新しいバージョンをインストールする
goenv install 1.25.3
# 新しいバージョンを現在のプロジェクトに設定する
goenv local 1.25.3
# 変更できたことを確認する
go version
# Goの現在の依存関係を解消する
go mod tidy
# 依存先を最新化: ./...はカレントディレクトリの配下全体
go get -u ./...
# 前のgo get -u ./...での更新を更に反映する
go mod tidy
```
- ビルドとテストを実行 
```bash
go build ./...
go test ./...
```
