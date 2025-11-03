package pkg

type Datum struct {
	Command     string
	Content     string
	Description string
}

type Data map[string][]Datum

func Commands() Data {
	return Data{
		"ノーマルモード": {
			{Command: "%", Content: "対照となるカッコ間を移動", Description: ""},
			{Command: "ci\"", Content: "\t\"\" で囲われたワードを削除", Description: "iの部分はaでもいい。a: 両端を含む, i: 両端を含まない"},
			{Command: "Shift + ^", Content: "行の先頭文字の頭まで移動", Description: ""},
			{Command: "f", Content: "同じ行から文字を探す(次のもの)", Description: ""},
			{Command: "F", Content: "同じ行から文字を探す(前のもの)", Description: ""},
			{Command: ";", Content: "次の検索結果にいく", Description: ""},
			{Command: ",", Content: "前の検索結果にいく", Description: "カスタムしているので現在は使えない"},
			{Command: "t", Content: "行末に向けて最初に見つかったもの", Description: ""},
			{Command: "T", Content: "行頭に向けて最初に見つかったもの", Description: ""},
			{Command: "{", Content: "次の文節に移動する", Description: "関数ごとなどで移動"},
			{Command: "}", Content: "前の文節に移動する", Description: "関数ごとなどで移動"},
			{Command: "'", Content: "前のマークの場所に移動", Description: ""},
			{Command: "ct=", Content: "=までの範囲を削除する(=を含まない)", Description: ""},
			{Command: "cf=", Content: "=までの範囲を削除する(=を含む)", Description: "const XXX = を return などに変えたい時に使う"},
			{Command: "gi", Content: "最後にinsertモードを終了した場所に戻ってinsertモードに切り替える", Description: ""},
			{Command: "ctrl + f", Content: "次へページ送り", Description: ""},
			{Command: "ctrl + b", Content: "前へページ送り", Description: ""},
			{Command: "zz", Content: "カーソル部が真ん中にいくように調整", Description: ""},
			{Command: "zt", Content: "カーソル部が上にいくように調整", Description: ""},
			{Command: "zb", Content: "カーソル部が下にいくように調整", Description: ""},
			{Command: "m{任意のキー}", Content: "任意のキーにマークをつけることができる", Description: ""},
			{Command: "`{任意のキー}", Content: "マークの位置に飛ぶことができる", Description: ""},
			{Command: "ctrl + o", Content: "直前のジャンプ前の位置に戻る", Description: ""},
			{Command: "ctrl + ]", Content: "呼び出し元にジャンプ/呼び出し先にジャンプ", Description: ""},
		},
		"ブックマーク(マーク)": {
			{Command: "m{a-z}", Content: "ローカルマーク(現在のファイル内)を設定", Description: "小文字のマークはファイル内でのみ有効"},
			{Command: "m{A-Z}", Content: "グローバルマーク(ファイル間)を設定", Description: "大文字のマークはファイルをまたいで有効"},
			{Command: "`{a-zA-Z}", Content: "マークした位置の正確な場所(行と列)に移動", Description: "バッククォートを使用"},
			{Command: "'{a-zA-Z}", Content: "マークした行の先頭に移動", Description: "シングルクォートを使用"},
			{Command: ":marks", Content: "設定されているマークの一覧を表示", Description: ""},
			{Command: ":delmarks {marks}", Content: "指定したマークを削除", Description: "例: :delmarks a b c"},
			{Command: ":delmarks!", Content: "すべてのローカルマーク(a-z)を削除", Description: ""},
			{Command: "``", Content: "最後にジャンプした位置に戻る", Description: "バッククォート2つ"},
			{Command: "''", Content: "最後にジャンプした行の先頭に戻る", Description: "シングルクォート2つ"},
			{Command: "`.", Content: "最後に編集した位置に移動", Description: ""},
			{Command: "`^", Content: "最後にInsertモードを抜けた位置に移動", Description: ""},
			{Command: "`[", Content: "最後にヤンクまたは変更したテキストの開始位置", Description: ""},
			{Command: "`]", Content: "最後にヤンクまたは変更したテキストの終了位置", Description: ""},
		},
		"コマンドモード": {
			{Command: "/{word}", Content: "n で次へ進み、N で 1 つ前に戻る", Description: "検索を行う"},
		},
		"ビジュアルブロックモード": {
			{Command: "ctrl-v -> I -> Esc", Content: "同じ箇所に同じ文字列を追加する", Description: ""},
			{Command: "ctrl-a / ctrl-x", Content: "数字のインクリメント/デクリメント", Description: ""},
			{Command: "g<ctrl-a> / g<ctrl-x>", Content: "数字の「連番で」インクリメント/デクリメント", Description: ""},
			{Command: "o", Content: "反対方向にブロックを伸ばすことができる", Description: ""},
		},
		"ファイル一覧": {
			{Command: "-", Content: "一つ上のディレクトリに移動", Description: ""},
			{Command: "u ←→ U", Content: "一つ前のディレクトリに戻る", Description: ""},
			{Command: "ctrl + w → s", Content: "水平分割", Description: ""},
			{Command: "ctrl + w → v", Content: "垂直分割", Description: ""},
			{Command: "ctrl + w → + h j k l\t", Content: "ブラウザ移動", Description: ""},
			{Command: "ctrl + w → + x or crtl + x\t", Content: "次のウィンドウと入れ替える", Description: ""},
			{Command: "ctrl + w → + r or crtl + r", Content: "ウィンドウ一を下 / 右へ回転させる", Description: ""},
			{Command: "ctrl + w → + R or crtl + R\t", Content: "ウィンドウ一下 / 右へ回転させる", Description: ""},
			{Command: "ctrl + w → K", Content: "ウィンドウを最上段に移動し、幅を最大にする", Description: ""},
			{Command: "ctrl + w → J", Content: "ウィンドウを最下段に移動し、幅を最大にする", Description: ""},
			{Command: "ctrl + w → H", Content: "ウィンドウを最左段に移動し、幅を最大にする", Description: ""},
			{Command: "ctrl + w → L", Content: "ウィンドウを最右段に移動し、幅を最大にする", Description: ""},
			{Command: "ctrl + w → T", Content: "ウィンドウを新しいタブページへ移動する", Description: ""},
		},
		"LSP補完": {
			{Command: "Ctl + y", Content: "選択中の補完を選択して適用", Description: ""},
		},
		"vim-surround": {
			{Command: "d s [e]", Content: "eを削除", Description: "delete surround"},
			{Command: "c s [e] [d]", Content: "eをd に変更", Description: "change surround"},
			{Command: "y s [mo] [d]", Content: "mo(モーション)を使って、dで囲む", Description: "you surround"},
			{Command: "S [d]", Content: "dで囲む（注： ヴィジュアルモードの時のみ）", Description: "Sorround(注: Sは大文字)"},
		},
	}
}

func Header() []string {
	return []string{"Title", "Command", "Description"}
}
