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
			{Command: "ci\"", Content: "\t\"\" で囲われたワードを削除", Description: ""},
			{Command: "Shift + ^", Content: "行の先頭文字の頭まで移動", Description: ""},
			{Command: "f", Content: "同じ行から文字を探す(次のもの)", Description: ""},
			{Command: "F", Content: "同じ行から文字を探す(前のもの)", Description: ""},
			{Command: ";", Content: "次の検索結果にいく", Description: ""},
			{Command: ",", Content: "前の検索結果にいく", Description: ""},
			{Command: "t", Content: "行末に向けて最初に見つかったもの", Description: ""},
			{Command: "T", Content: "行頭に向けて最初に見つかったもの", Description: ""},
		},
		"コマンドモード": {
			{Command: "/{word}", Content: "n で次へ進み、N で 1 つ前に戻る", Description: "検索を行う"},
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
