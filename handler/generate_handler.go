package handler

import (
	"bytes"
	"cc-generate-course-service/helper"
	"cc-generate-course-service/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GenerateHandler struct{}

func NewGenerateHandler() *GenerateHandler {
	return &GenerateHandler{}
}

func GenerateCourseHandler(c *fiber.Ctx) error {
	url := "http://localhost:8081/api/v1/course/create"

	token := c.Get("Authorization")
	token = token[len("Bearer "):]

	var prompt model.Prompt
	if err := c.BodyParser(&prompt); err != nil {
		return helper.Response(c, 400, "Error Parsing the Body", nil)
	}

	if prompt.Prompt == "" {
		return helper.Response(c, 400, "Prompt is Empty", nil)
	}

	raw, err := helper.GenerateCourse(prompt.Prompt)
	if err != nil {
		return helper.Response(c, 400, "Failed to Generate Course", nil)
	}

	requestBody := map[string]interface{}{
		"course": raw,
		"prompt": prompt.Prompt,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp.StatusCode :", resp.StatusCode)
		return fmt.Errorf("request failed with status: %v", resp.Status)
	}
	fmt.Println("resp.StatusCode :", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body :", err)
		return err
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		// Handle the case where "data" is not present or is not a map
		return errors.New("data is not present or is not a map")
	}

	// Check if the "course_id" key is present in the "data" map
	courseID, ok := data["course_id"].(string)
	if !ok {
		// Handle the case where "course_id" is not present or is not a string
		return errors.New("course_id is not present or is not a string")
	}

	// Use the courseID
	fmt.Println("course_id:", courseID)

	// Return the response
	return helper.Response(c, 200, "Success", fiber.Map{
		"course_id": courseID,
		"message":   "Course Generated Successfully",
	})
}
