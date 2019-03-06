package params

type FaceDetect struct {
	ReturnFaceId         *bool   `url:"returnFaceId,omitempty"`
	ReturnFaceLandmarks  *bool   `url:"returnFaceLandmarks,omitempty"`
	ReturnFaceAttributes *string `url:"returnFaceAttributes,omitempty"`
}

type PersonGroupPersonCreateResp struct {
	PersonID string `json:"personId"`
}

type XY struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type FaceDetectResp struct {
	FaceID        string `json:"faceId"`
	FaceRectangle struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Left   int `json:"left"`
		Top    int `json:"top"`
	} `json:"faceRectangle"`
	FaceLandmarks  map[string]XY `json:"faceLandmarks"`
	FaceAttributes struct {
		Age        float64 `json:"age"`
		Gender     string  `json:"gender"`
		Smile      float64 `json:"smile"`
		FacialHair struct {
			Moustache float64 `json:"moustache"`
			Beard     float64 `json:"beard"`
			Sideburns float64 `json:"sideburns"`
		} `json:"facialHair"`
		Glasses  string `json:"glasses"`
		HeadPose struct {
			Roll  float64 `json:"roll"`
			Yaw   int     `json:"yaw"`
			Pitch int     `json:"pitch"`
		} `json:"headPose"`
		Emotion struct {
			Anger     float64 `json:"anger"`
			Contempt  int     `json:"contempt"`
			Disgust   float64 `json:"disgust"`
			Fear      float64 `json:"fear"`
			Happiness float64 `json:"happiness"`
			Neutral   float64 `json:"neutral"`
			Sadness   int     `json:"sadness"`
			Surprise  float64 `json:"surprise"`
		} `json:"emotion"`
		Hair struct {
			Bald      float64 `json:"bald"`
			Invisible bool    `json:"invisible"`
			HairColor []struct {
				Color      string  `json:"color"`
				Confidence float64 `json:"confidence"`
			} `json:"hairColor"`
		} `json:"hair"`
		Makeup struct {
			EyeMakeup bool `json:"eyeMakeup"`
			LipMakeup bool `json:"lipMakeup"`
		} `json:"makeup"`
		Occlusion struct {
			ForeheadOccluded bool `json:"foreheadOccluded"`
			EyeOccluded      bool `json:"eyeOccluded"`
			MouthOccluded    bool `json:"mouthOccluded"`
		} `json:"occlusion"`
		Accessories []struct {
			Type       string  `json:"type"`
			Confidence float64 `json:"confidence,omitempty"`
		} `json:"accessories"`
		Blur struct {
			BlurLevel string  `json:"blurLevel"`
			Value     float64 `json:"value"`
		} `json:"blur"`
		Exposure struct {
			ExposureLevel string  `json:"exposureLevel"`
			Value         float64 `json:"value"`
		} `json:"exposure"`
		Noise struct {
			NoiseLevel string  `json:"noiseLevel"`
			Value      float64 `json:"value"`
		} `json:"noise"`
	} `json:"faceAttributes"`
}

type FaceIdentifyResp struct {
	FaceID     string      `json:"faceId"`
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	PersonID   string  `json:"personId"`
	Confidence float64 `json:"confidence"`
}
