package webui

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	_ "embed"

	"github.com/gobuffalo/packr/v2"
)

//go:embed api/version
var version string

var Port = 0
var Host = "https://squiz.mrcyjanek.net"
var httpClient = http.Client{}

// Response -
type Response struct {
	OK      bool   `json:"ok"`      // true,
	Message string `json:"message"` // This one is for errors
	Version string `json:"version"` // "v2",
	Answers []struct {
		ID   string `json:"_id"`  // "5ebce979396b57001b9be31c",
		Type string `json:"type"` // "MSQ",
		//Ver       int    `json:"ver"`  // 2,
		//Published bool   `json:"published"`
		Structure struct {
			Settings struct {
				HasCorrectAnswer bool `json:"hasCorrectAnswer"` // true
			} `json:"settings"`
			Kind  string `json:"kind"` // "MSQ",
			Query struct {
				Math struct {
					Latex    []string `json:"latex"`
					Template string   `json:"template"` // null
				} `json:"math"`
				Type    string `json:"type"`    // null,
				HasMath bool   `json:"hasMath"` // false,
				Text    string `json:"text"`    // "<p>2<\/p>",
				Media   []struct {
					Type string `json:"type"`
					URL  string `json:"url"` // "https:\/\/quizizz.com\/media\/resource\/gs\/quizizz-media\/quizzes\/97f821e6-6350-47cc-8898-91db0d29ba07",
					Meta struct {
						Width   int    `json:"width"`   // 996,
						Height  int    `json:"height"`  // 352,
						Text    string `json:"text"`    // null,
						BGColor string `json:"bgColor"` // null
					} `json:"meta"`
				} `json:"media"`
			} `json:"query"`
			Options []struct {
				Math struct {
					Latex    []string `json:"latex"`
					Template string   `json:"template"` // null
				} `json:"math"`
				Type    string `json:"type"`    // null,
				HasMath bool   `json:"hasMath"` // false,
				Text    string `json:"text"`    // "<p>2<\/p>",
				Media   []struct {
					Type string `json:"type"`
					URL  string `json:"url"` // "https:\/\/quizizz.com\/media\/resource\/gs\/quizizz-media\/quizzes\/97f821e6-6350-47cc-8898-91db0d29ba07",
					Meta struct {
						Width   int    `json:"width"`   // 996,
						Height  int    `json:"height"`  // 352,
						Text    string `json:"text"`    // null,
						BGColor string `json:"bgColor"` // null
					} `json:"meta"`
				} `json:"media"`
			} `json:"options"`
			HasMath bool `json:"hasMath"` // false
		} `json:"structure"`
		//CreatedAt string `json:"createdAt"` // "2020-05-14T06:47:21.150Z",
		//Updated   string `json:"updated"`   // "2020-11-24T11:59:41.717Z",
		//DashDashV int    `json:"__v"`       // 0,
		//Time      int    `json:"time"`      // 30000,
		//State     string `json:"state"`     // "inactive",
		//Attempt   int    `json:"attempt"`   // 0,
		//Cause     string `json:"cause"`     // "",
		Answer struct {
			Answer  interface{} `json:"answer"`
			Options []struct {
				Math struct {
					Latex    []string `json:"latex"`    // [],
					Template string   `json:"template"` // null
				} `json:"math"`
				Type    string `json:"type"`    // "text",
				HasMath bool   `json:"hasMath"` // false,
				Text    string `json:"text"`    // "2",
				Media   []struct {
					Type string `json:"type"`
					URL  string `json:"url"` // "https:\/\/quizizz.com\/media\/resource\/gs\/quizizz-media\/quizzes\/97f821e6-6350-47cc-8898-91db0d29ba07",
					Meta struct {
						Width   int    `json:"width"`   // 996,
						Height  int    `json:"height"`  // 352,
						Text    string `json:"text"`    // null,
						BGColor string `json:"bgColor"` // null
					} `json:"meta"`
				} `json:"media"`
			} `json:"options"`
		} `json:"answer"`
	} `json:"answers"`
}

var answers Response

// Start the webui
func Start() {
	if Port == 0 {
		Port = 2000 + rand.Intn(10000)
	}
	Port = 15932 // Sorry for hardcoding
	html := packr.New("webui", "./html")
	http.Handle("/", http.FileServer(html))
	vappend := flag.String("version", "", "Append something to version :)")
	flag.Parse()
	version = version + " " + *vappend
	http.HandleFunc("/api/hack", apiHack)
	http.HandleFunc("/api/version", apiVersion)
	//http.HandleFunc("/api/answers", apiAnswers)
	//http.HandleFunc("/answers", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Add("Content-Type", "text/html;charset=utf-8")
	//	fmt.Fprintln(w, "<!DOCTYPE html>\n<html>\n<head>\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">")
	//	fmt.Fprintln(w, "<style>body { font-size: 150% }</style>")
	//	fmt.Fprintln(w, `<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/katex.min.css" integrity="sha384-AfEj0r4/OFrOo5t7NnNe46zW/tFgW6x/bCJG8FqQCEo3+Aro6EYUG4+cU+KJWu/X" crossorigin="anonymous">
	//	<!-- The loading of KaTeX is deferred to speed up page rendering -->
	//	<script defer src="https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/katex.min.js" integrity="sha384-g7c+Jr9ZivxKLnZTDUhnkOnsh30B4H0rpLUpJ4jAIKs4fnJI+sEnkvrMWph2EDg4" crossorigin="anonymous"></script>
	//	<!-- To automatically render math in text elements, include the auto-render extension: -->
	//	<script defer src="https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/contrib/auto-render.min.js" integrity="sha384-mll67QQFJfxn0IYznZYonOWZ644AWYC+Pt2cHqMaRhXVrursRwvLnLaebdGIlYNa" crossorigin="anonymous"
	//		onload="renderMathInElement(document.body);"></script>`)
	//	fmt.Fprintln(w, "<body><div style=\"margin: auto; max-width:650px;\">")
	//	for q := range answers.Answers {
	//		answ := answers.Answers[q]
	//		//var correct []float32
	//		//switch reflect.TypeOf(answ.Answer.Answer).String() {
	//		//case "[]interface {}":
	//		//	correct = append(correct, reflect.(answ.Answer.Answer))
	//		//	break
	//		//}
	//		fmt.Fprintln(w, "<hr />"+answ.Structure.Query.Text+strings.Join(answ.Structure.Query.Math.Latex, "<br />\n")+"<br />\n")
	//		for p := range answ.Structure.Options {
	//			if strings.Contains(fmt.Sprintf("%v", answ.Answer.Answer), strconv.Itoa(p)) {
	//				fmt.Fprintln(w, "<span style=\"color: green\">")
	//			}
	//			opts := answ.Structure.Options[p]
	//			answString := opts.Text + strings.Join(opts.Math.Latex, "\n<br />")
	//			answString = strings.ReplaceAll(answString, "<p", "<span")
	//			answString = strings.ReplaceAll(answString, "p/>", "span/>")
	//			answString = strings.ReplaceAll(answString, "/p>", "/span>")
	//			answString = strings.ReplaceAll(answString, "/ p>", "/ span>")
	//			fmt.Fprintln(w, strconv.Itoa(p)+") "+answString+"<br />\n")
	//			for m := range opts.Media {
	//				media := opts.Media[m]
	//				fmt.Fprintln(w, "<img src=\""+media.URL+"\"></img><br />\n")
	//			}
	//			if strings.Contains(fmt.Sprintf("%v", answ.Answer.Answer), strconv.Itoa(p)) {
	//				fmt.Fprintln(w, "</span>")
	//			}
	//		}
	//		for m := range answ.Structure.Query.Media {
	//			media := answ.Structure.Query.Media[m]
	//			fmt.Fprintln(w, "<img src=\""+media.URL+"\"></img><br />\n")
	//		}
	//		fmt.Fprintf(w, "Answer: "+fmt.Sprintf("%v", answ.Answer.Answer)+"<br />\n")
	//		for j := range answ.Answer.Options {
	//			opts := answ.Answer.Options[j]
	//			fmt.Fprintln(w, strconv.Itoa(j)+") "+opts.Text+strings.Join(opts.Math.Latex, "\n<br />")+"<br />\n")
	//			for m := range opts.Media {
	//				media := opts.Media[m]
	//				fmt.Fprintln(w, "<img src=\""+media.URL+"\"></img><br />\n")
	//			}
	//		}
	//		fmt.Fprintf(w, "")
	//	}
	//	fmt.Fprintln(w, "<hr /> Owned by Czarek Nakamoto | <a href=\"https://mrcyjanek.net\">mrcyjanek.net</a>")
	//	fmt.Fprintln(w, "</div></body></html>")
	//	//fmt.Fprintln(w, string(out))
	//})
	//http.HandleFunc("/api/ping", apiPing)
	go http.ListenAndServe(":"+strconv.Itoa(Port), nil)
	fmt.Println("[webui][Start] Listening on 127.0.0.1:" + strconv.Itoa(Port))
}

func apiHack(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	realHackOk(req.FormValue("pin"), req.FormValue("key"))
	r := answers
	out, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(out)
}

func apiVersion(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(version))
}
