package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const IMAGE_GENERATE_URL = "https://api.openai.com/v1/images/generations"
const OPEN_AI_API_KEY = "Bearer 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"

type imageGenerateRequest struct {
	Text string `json:"prompt"`
	N    int    `json:"n"`
	Size string `json:"size"`
}

type imageGenerateResponse struct {
	Created int                 `json:"created"`
	Data    []map[string]string `json:"data"`
}

func GetImageByText(ginContext *gin.Context) {
	//The text to generate the image from
	text := ginContext.Param("text")

	//setting the image request data
	data := imageGenerateRequest{
		text,
		1,
		"1024x1024",
	}

	//seralizing the data into
	out, err := json.Marshal(data)
	if err != nil {
		ginContext.JSON(400, gin.H{"error": err})
		return
	}
	payloadBuf := bytes.NewBuffer(out)
	//creating new POST request with data as the body
	req, err := http.NewRequest("POST", IMAGE_GENERATE_URL, payloadBuf)
	if err != nil {
		ginContext.JSON(req.Response.StatusCode, gin.H{"error": err})
		return
	}

	//setting request headers
	req.Header.Set("Authorization", OPEN_AI_API_KEY)
	req.Header.Set("Content-Type", "application/json")

	//sending the POST request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		ginContext.JSON(response.StatusCode, gin.H{"error": err})
		return
	}
	defer response.Body.Close()

	if err != nil {
		ginContext.JSON(response.StatusCode, gin.H{"error": err})
		return
	}

	//reading json data into a imageGenerateResponse struct
	imageUrl := imageGenerateResponse{}
	body, error := io.ReadAll(response.Body)
	if error != nil {
		ginContext.JSON(response.StatusCode, gin.H{"error": err})
		return
	}

	//deserializing the json into imageUrl
	if err := json.Unmarshal(body, &imageUrl); err != nil {
		ginContext.JSON(response.StatusCode, gin.H{"error": err})
		return
	}

	//showing imageUrl data as json
	ginContext.PureJSON(200, imageUrl.Data)
}
