package face

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/eternnoir/go-azure-face/params"

	"github.com/levigross/grequests"
)

const apiPersonGroups = "persongroups"

const (
	headerOctetStream = "application/octet-stream"
	headerContentType = "Content-Type"
)

// Face is main struct for azure face api.
type Face struct {
	apiURL string
	apiKey string
}

// New return new Face api instance.
func New(apiurl, apikey string) *Face {
	return &Face{apiurl, apikey}
}

func (f *Face) FaceDetect(returnFaceId, returnFaceLandmarks *bool, returnFaceAttributes, url *string, data []byte) ([]params.FaceDetectResp, error) {
	if url == nil && data == nil {
		return nil, errors.New("url or data must have one")
	}

	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Ocp-Apim-Subscription-Key": f.apiKey},
	}

	if url != nil {
		data := map[string]interface{}{}
		data["url"] = *url
		ro.JSON = data
	}
	if data != nil {
		ro.Headers[headerContentType] = headerOctetStream
		ro.RequestBody = bytes.NewBuffer(data)
	}
	resp, err := grequests.Post(fmt.Sprintf("%s/%s", f.apiURL, "detect"), ro)
	if err != nil {
		return nil, err
	}
	var ret []params.FaceDetectResp
	_, err = checkResp(resp, &ret)
	return ret, err
}

func (f *Face) FaceIdentify(faceIds []string, personGroupId, largePersonGroupId string, maxNumOfCandidatesReturned, confidenceThreshold *int) ([]params.FaceIdentifyResp, error) {
	data := map[string]interface{}{}
	data["faceIds"] = faceIds
	if personGroupId != "" {
		data["personGroupId"] = personGroupId
	}
	if largePersonGroupId != "" {
		data["largePersonGroupId"] = largePersonGroupId
	}
	if maxNumOfCandidatesReturned != nil {
		data["maxNumOfCandidatesReturned"] = *maxNumOfCandidatesReturned
	}
	if confidenceThreshold != nil {
		data["confidenceThreshold "] = confidenceThreshold
	}
	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Ocp-Apim-Subscription-Key": f.apiKey},
		JSON:    data,
	}
	resp, err := grequests.Post(fmt.Sprintf("%s/%s", f.apiURL, "identify"), ro)
	if err != nil {
		return nil, err
	}
	var ret []params.FaceIdentifyResp
	_, err = checkResp(resp, &ret)
	return ret, err
}

func (f *Face) PersonGroupCreate(personGroupId, name string, userData *string) error {
	data := map[string]interface{}{}
	data["name"] = name
	if userData != nil {
		data["userData"] = userData
	}

	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Ocp-Apim-Subscription-Key": f.apiKey},
		JSON:    data,
	}
	resp, err := grequests.Put(fmt.Sprintf("%s/%s/%s", f.apiURL, apiPersonGroups, personGroupId), ro)
	if err != nil {
		return err
	}
	_, err = checkResp(resp, nil)
	return err
}

func (f *Face) PersonGroupTrain(personGroupId string) error {
	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Ocp-Apim-Subscription-Key": f.apiKey},
	}
	resp, err := grequests.Post(fmt.Sprintf("%s/%s/%s/train", f.apiURL, apiPersonGroups, personGroupId), ro)
	if err != nil {
		return err
	}
	_, err = checkResp(resp, nil)
	return err
}

func (f *Face) PersonGroupPersonCreate(personGroupId, name string, userData *string) (*params.PersonGroupPersonCreateResp, error) {
	data := map[string]interface{}{}
	data["name"] = name
	if userData != nil {
		data["userData"] = userData
	}

	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Ocp-Apim-Subscription-Key": f.apiKey},
		JSON:    data,
	}
	resp, err := grequests.Post(fmt.Sprintf("%s/%s/%s/%s", f.apiURL, apiPersonGroups, personGroupId, "persons"), ro)
	if err != nil {
		return nil, err
	}
	var ret params.PersonGroupPersonCreateResp
	_, err = checkResp(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (f *Face) PersonGroupPersonAddFace(personGroupId, personId string, userData, targetFace, url *string, data []byte) error {
	if url == nil && data == nil {
		return errors.New("url or data must have one")
	}

	query := map[string]string{}
	if userData != nil {
		query["userData"] = *userData
	}
	if targetFace != nil {
		query["targetFace"] = *targetFace
	}

	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Ocp-Apim-Subscription-Key": f.apiKey},
	}

	if url != nil {
		data := map[string]interface{}{}
		data["url"] = *url
		ro.JSON = data
	}
	if data != nil {
		ro.Headers[headerContentType] = headerOctetStream
		ro.RequestBody = bytes.NewBuffer(data)
	}
	resp, err := grequests.Post(fmt.Sprintf("%s/%s/%s/%s/%s/persistedFaces", f.apiURL, apiPersonGroups, personGroupId, "persons", personId), ro)
	if err != nil {
		return err
	}
	_, err = checkResp(resp, nil)
	return err
}

func checkResp(reqresp *grequests.Response, resp interface{}) (interface{}, error) {
	if !reqresp.Ok {
		var apierr ApiError
		if err := json.Unmarshal(reqresp.Bytes(), &apierr); err != nil {
			return nil, err
		}
		return nil, apierr
	}
	if resp == nil {
		return nil, nil
	}
	if err := json.Unmarshal(reqresp.Bytes(), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
