package pkg

type Datum struct {
	Command     string
	Content     string
	Description string
}

type Data map[string][]Datum

func Commands() Data {
	return Data{
		"vimrc": {
			{Command: "key1/title1", Content: "key1/content1", Description: "key1/description1"},
			{Command: "key1/title2", Content: "key1/content2", Description: "key1/description2"},
		},
		"suround": {
			{Command: "key2/title1", Content: "key2/content1", Description: "key2/description1"},
			{Command: "key2/title2", Content: "key2/content2", Description: "key2/description2"},
		},
	}
}

func Header() []string {
	return []string{"Title", "Command", "Description"}
}
