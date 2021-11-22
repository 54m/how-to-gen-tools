# how-to-gen-tools
パートナー向け内製ツールの説明会(api_gen, firestore-repo)アジェンダ

# api_genって何
- server generator
  - APIのRoutingやControllerを階層から自動的に判定し、生成する。
  - 自動生成されたControllerに対して実際の処理を書いていく。
  - 基本的に生成物は既存のファイルを上書きする。(Controllerはすでにファイルが存在した場合は上書きしない。)

- client generator
  - `api_gen server` 向けの用意したディレクトリに対して実行する。
  - ライブラリは今いるディレクトリに対して生成される。
  - TypeScript+fetchを利用したライブラリが生成される。

# firestore-repoって何
Cloud Firestoreで利用されるコード(CRUD & Search)を自動生成する。

# デモ

---

# ブランチの説明
- main
  - ディレクトリ構成の説明
- api_gen
  - api_gen生成前
- api_gen-generated
  - api_gen生成後
- firestore-repo
  - firestore-repo生成前
- firestore-repo-generated
  - firestore-repo生成後
- completed
  - 完成系

# 質疑応答

---

## License
- Under the [MIT License](./LICENSE)
- Copyright (C) 2021 54m
