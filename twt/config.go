package twt

type (
	MirrorOption struct {
		Username string
		Active   struct {
			Period        string
			CheckInterval string
		}
		InActive struct {
			Period        string
			CheckInterval string
		}
	}
)
