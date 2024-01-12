package JsonMapper

// FindX: safety find
// if null, return default value

func (j *JsonMap) FindBool(k string, defaultValue bool) bool {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return tmp.(bool)
	}
}

func (j *JsonMap) FindString(k string, defaultValue string) string {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return tmp.(string)
	}
}

func (j *JsonMap) FindInt(k string, defaultValue int) int {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return int(tmp.(float64))
	}
}

func (j *JsonMap) FindInt32(k string, defaultValue int32) int32 {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return int32(tmp.(float64))
	}
}

func (j *JsonMap) FindInt64(k string, defaultValue int64) int64 {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return int64(tmp.(float64))
	}
}

func (j *JsonMap) FindFloat32(k string, defaultValue float32) float32 {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return float32(tmp.(float64))
	}
}

func (j *JsonMap) FindFloat64(k string, defaultValue float64) float64 {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return tmp.(float64)
	}
}

func (j *JsonMap) FindMap(k string, defaultValue map[string]interface{}) map[string]interface{} {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return tmp.(map[string]interface{})
	}
}

func (j *JsonMap) FindSilce(k string, defaultValue []interface{}) []interface{} {
	tmp := j.Find(k)
	if tmp == nil {
		return defaultValue
	} else {
		return tmp.([]interface{})
	}
}
