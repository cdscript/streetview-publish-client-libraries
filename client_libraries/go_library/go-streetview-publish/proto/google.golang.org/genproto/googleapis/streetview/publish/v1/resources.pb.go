// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/streetview/publish/v1/resources.proto

package publish

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"
import google_type "google.golang.org/genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Upload reference for media files.
type UploadRef struct {
	// Required. An upload reference should be unique for each user. It follows
	// the form (depending on which RPC is called):
	// <ul><li>
	// Photos:
	// "https://streetviewpublish.googleapis.com/media/user/<account_id>/photo/<upload_reference>"
	// </li><li>
	// Photo sequences (in the form of videos):
	// "https://streetviewpublish.googleapis.com/media/user/<account_id>/video/<upload_reference>"
	// </li></ul>
	UploadUrl string `protobuf:"bytes,1,opt,name=upload_url,json=uploadUrl" json:"upload_url,omitempty"`
}

func (m *UploadRef) Reset()                    { *m = UploadRef{} }
func (m *UploadRef) String() string            { return proto.CompactTextString(m) }
func (*UploadRef) ProtoMessage()               {}
func (*UploadRef) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *UploadRef) GetUploadUrl() string {
	if m != nil {
		return m.UploadUrl
	}
	return ""
}

// Identifier for a photo.
type PhotoId struct {
	// Required. A base64 encoded identifier.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *PhotoId) Reset()                    { *m = PhotoId{} }
func (m *PhotoId) String() string            { return proto.CompactTextString(m) }
func (*PhotoId) ProtoMessage()               {}
func (*PhotoId) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PhotoId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Raw pose measurement for an entity.
type Pose struct {
	// Latitude and longitude pair of the pose, as explained here:
	// https://developers.google.com/my-business/reference/rest/Shared.Types/LatLng
	// When creating a photo, if the latitude and longitude pair are not provided
	// here, the geolocation from the exif header will be used.
	// If the latitude and longitude pair is not provided and cannot be found in
	// the exif header, the create photo process will fail.
	LatLngPair *google_type.LatLng `protobuf:"bytes,1,opt,name=lat_lng_pair,json=latLngPair" json:"lat_lng_pair,omitempty"`
	// Altitude of the pose in meters above ground level (as defined by WGS84).
	// NaN indicates an unmeasured quantity.
	Altitude float64 `protobuf:"fixed64,2,opt,name=altitude" json:"altitude,omitempty"`
	// Compass heading, measured at the center of the photo in degrees clockwise
	// from North. Value must be >=0 and <360.
	// NaN indicates an unmeasured quantity.
	Heading float64 `protobuf:"fixed64,3,opt,name=heading" json:"heading,omitempty"`
	// Pitch, measured at the center of the photo in degrees. Value must be >=-90
	// and <= 90. A value of -90 means looking directly down, and a value of 90
	// means looking directly up.
	// NaN indicates an unmeasured quantity.
	Pitch float64 `protobuf:"fixed64,4,opt,name=pitch" json:"pitch,omitempty"`
	// Roll, measured in degrees. Value must be >-180 and <=180. A value of 0
	// means level with the horizon.
	// NaN indicates an unmeasured quantity.
	Roll float64 `protobuf:"fixed64,5,opt,name=roll" json:"roll,omitempty"`
}

func (m *Pose) Reset()                    { *m = Pose{} }
func (m *Pose) String() string            { return proto.CompactTextString(m) }
func (*Pose) ProtoMessage()               {}
func (*Pose) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Pose) GetLatLngPair() *google_type.LatLng {
	if m != nil {
		return m.LatLngPair
	}
	return nil
}

func (m *Pose) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Pose) GetHeading() float64 {
	if m != nil {
		return m.Heading
	}
	return 0
}

func (m *Pose) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *Pose) GetRoll() float64 {
	if m != nil {
		return m.Roll
	}
	return 0
}

// Place metadata for an entity.
type Place struct {
	// Required. Place identifier, as described in
	// https://developers.google.com/places/place-id.
	PlaceId string `protobuf:"bytes,1,opt,name=place_id,json=placeId" json:"place_id,omitempty"`
}

func (m *Place) Reset()                    { *m = Place{} }
func (m *Place) String() string            { return proto.CompactTextString(m) }
func (*Place) ProtoMessage()               {}
func (*Place) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Place) GetPlaceId() string {
	if m != nil {
		return m.PlaceId
	}
	return ""
}

// A connection is the link from a source photo to a destination photo.
type Connection struct {
	// Required. The destination of the connection from the containing photo to
	// another photo.
	Target *PhotoId `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
}

func (m *Connection) Reset()                    { *m = Connection{} }
func (m *Connection) String() string            { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()               {}
func (*Connection) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Connection) GetTarget() *PhotoId {
	if m != nil {
		return m.Target
	}
	return nil
}

// Photo is used to store 360 photos along with photo metadata.
type Photo struct {
	// Output only. Identifier for the photo, which is unique among all photos in
	// Google.
	PhotoId *PhotoId `protobuf:"bytes,1,opt,name=photo_id,json=photoId" json:"photo_id,omitempty"`
	// Required (when creating photo). Input only. The resource URL where the
	// photo bytes are uploaded to.
	UploadReference *UploadRef `protobuf:"bytes,2,opt,name=upload_reference,json=uploadReference" json:"upload_reference,omitempty"`
	// Output only. The download URL for the photo bytes. This field is set only
	// when the `view` parameter in a `GetPhotoRequest` is set to
	// `INCLUDE_DOWNLOAD_URL`.
	DownloadUrl string `protobuf:"bytes,3,opt,name=download_url,json=downloadUrl" json:"download_url,omitempty"`
	// Output only. The thumbnail URL for showing a preview of the given photo.
	ThumbnailUrl string `protobuf:"bytes,9,opt,name=thumbnail_url,json=thumbnailUrl" json:"thumbnail_url,omitempty"`
	// Output only. The share link for the photo.
	ShareLink string `protobuf:"bytes,11,opt,name=share_link,json=shareLink" json:"share_link,omitempty"`
	// Pose of the photo.
	Pose *Pose `protobuf:"bytes,4,opt,name=pose" json:"pose,omitempty"`
	// Connections to other photos. A connection represents the link from this
	// photo to another photo.
	Connections []*Connection `protobuf:"bytes,5,rep,name=connections" json:"connections,omitempty"`
	// Absolute time when the photo was captured.
	// When the photo has no exif timestamp, this is used to set a timestamp in
	// the photo metadata.
	CaptureTime *google_protobuf3.Timestamp `protobuf:"bytes,6,opt,name=capture_time,json=captureTime" json:"capture_time,omitempty"`
	// Places where this photo belongs.
	Places []*Place `protobuf:"bytes,7,rep,name=places" json:"places,omitempty"`
	// Output only. View count of the photo.
	ViewCount int64 `protobuf:"varint,10,opt,name=view_count,json=viewCount" json:"view_count,omitempty"`
}

func (m *Photo) Reset()                    { *m = Photo{} }
func (m *Photo) String() string            { return proto.CompactTextString(m) }
func (*Photo) ProtoMessage()               {}
func (*Photo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Photo) GetPhotoId() *PhotoId {
	if m != nil {
		return m.PhotoId
	}
	return nil
}

func (m *Photo) GetUploadReference() *UploadRef {
	if m != nil {
		return m.UploadReference
	}
	return nil
}

func (m *Photo) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *Photo) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

func (m *Photo) GetShareLink() string {
	if m != nil {
		return m.ShareLink
	}
	return ""
}

func (m *Photo) GetPose() *Pose {
	if m != nil {
		return m.Pose
	}
	return nil
}

func (m *Photo) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func (m *Photo) GetCaptureTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CaptureTime
	}
	return nil
}

func (m *Photo) GetPlaces() []*Place {
	if m != nil {
		return m.Places
	}
	return nil
}

func (m *Photo) GetViewCount() int64 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func init() {
	proto.RegisterType((*UploadRef)(nil), "google.streetview.publish.v1.UploadRef")
	proto.RegisterType((*PhotoId)(nil), "google.streetview.publish.v1.PhotoId")
	proto.RegisterType((*Pose)(nil), "google.streetview.publish.v1.Pose")
	proto.RegisterType((*Place)(nil), "google.streetview.publish.v1.Place")
	proto.RegisterType((*Connection)(nil), "google.streetview.publish.v1.Connection")
	proto.RegisterType((*Photo)(nil), "google.streetview.publish.v1.Photo")
}

func init() { proto.RegisterFile("google/streetview/publish/v1/resources.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0xd5, 0x76, 0xb7, 0xbb, 0xdd, 0xc9, 0xf2, 0x21, 0xc3, 0x21, 0x5d, 0x81, 0x80, 0x54, 0x88,
	0x0a, 0xa1, 0x44, 0x2d, 0x82, 0x4b, 0x85, 0x84, 0xda, 0x53, 0xa1, 0x87, 0x95, 0xa1, 0x1c, 0xb8,
	0x44, 0xde, 0xc4, 0x4d, 0x2c, 0xbc, 0xb6, 0xe5, 0x38, 0xad, 0xf8, 0x0d, 0xfc, 0x06, 0x4e, 0xfc,
	0x51, 0x6c, 0xc7, 0x49, 0x39, 0x54, 0x2d, 0xa7, 0xcc, 0x3c, 0xbf, 0x37, 0xf6, 0xcc, 0xb3, 0x03,
	0x6f, 0x2a, 0x29, 0x2b, 0x4e, 0xb3, 0xc6, 0x68, 0x4a, 0xcd, 0x25, 0xa3, 0x57, 0x99, 0x6a, 0xd7,
	0x9c, 0x35, 0x75, 0x76, 0x79, 0x90, 0x69, 0xda, 0xc8, 0x56, 0x17, 0xb4, 0x49, 0x95, 0x96, 0x46,
	0xa2, 0x27, 0x1d, 0x3b, 0xbd, 0x66, 0xa7, 0x81, 0x9d, 0x5e, 0x1e, 0x2c, 0xc3, 0x6a, 0x46, 0x14,
	0xcb, 0x88, 0x10, 0xd2, 0x10, 0xc3, 0xa4, 0x08, 0xda, 0xe5, 0xb3, 0xb0, 0xea, 0xb3, 0x75, 0x7b,
	0x91, 0x19, 0xb6, 0xa1, 0x8d, 0x21, 0x1b, 0x15, 0x08, 0x71, 0x20, 0x98, 0x9f, 0x8a, 0x66, 0x9c,
	0x18, 0x2e, 0xaa, 0x6e, 0x25, 0x79, 0x0d, 0xf3, 0x73, 0xc5, 0x25, 0x29, 0x31, 0xbd, 0x40, 0x4f,
	0x01, 0x5a, 0x9f, 0xe4, 0xad, 0xe6, 0xf1, 0xe8, 0xf9, 0x68, 0x7f, 0x8e, 0xe7, 0x1d, 0x72, 0xae,
	0x79, 0xb2, 0x0b, 0xb3, 0x55, 0x6d, 0x45, 0xa7, 0x25, 0xba, 0x0f, 0x5b, 0xac, 0x0c, 0x0c, 0x1b,
	0x25, 0xbf, 0x47, 0x30, 0x59, 0xc9, 0x86, 0xa2, 0x77, 0xb0, 0xb0, 0xf5, 0x73, 0xbb, 0x41, 0xae,
	0x08, 0xd3, 0x9e, 0x12, 0x1d, 0x3e, 0x4a, 0x43, 0x77, 0xee, 0x00, 0xe9, 0x19, 0x31, 0x67, 0xa2,
	0xc2, 0xc0, 0xfd, 0x77, 0x65, 0x69, 0x68, 0x09, 0x3b, 0x84, 0x1b, 0x66, 0xda, 0x92, 0xc6, 0x5b,
	0x56, 0x32, 0xc2, 0x43, 0x8e, 0x62, 0x98, 0xd5, 0x94, 0x94, 0x4c, 0x54, 0xf1, 0xd8, 0x2f, 0xf5,
	0x29, 0x7a, 0x0c, 0xdb, 0x8a, 0x99, 0xa2, 0x8e, 0x27, 0x1e, 0xef, 0x12, 0x84, 0x60, 0xa2, 0x25,
	0xe7, 0xf1, 0xb6, 0x07, 0x7d, 0x9c, 0x24, 0xb0, 0xbd, 0xe2, 0xa4, 0xa0, 0x68, 0x17, 0x76, 0x94,
	0x0b, 0xf2, 0xe1, 0xf8, 0x33, 0x9f, 0x9f, 0x96, 0xc9, 0x67, 0x80, 0x13, 0x29, 0x04, 0x2d, 0xdc,
	0x68, 0xd1, 0x07, 0x98, 0x1a, 0xa2, 0x2b, 0x6a, 0x42, 0x0b, 0x2f, 0xd3, 0xdb, 0x0c, 0x4a, 0xc3,
	0x60, 0x70, 0x10, 0x25, 0x7f, 0x26, 0x76, 0x47, 0x87, 0xa1, 0x8f, 0x76, 0x47, 0x17, 0xf4, 0x3b,
	0xfe, 0x77, 0xa9, 0x99, 0x0a, 0xc3, 0xc6, 0xf0, 0x30, 0xd8, 0xa2, 0xe9, 0x05, 0xd5, 0x54, 0x14,
	0xdd, 0x90, 0xa2, 0xc3, 0x57, 0xb7, 0x57, 0x1a, 0x9c, 0xc5, 0x0f, 0xda, 0x3e, 0xec, 0xf4, 0xe8,
	0x05, 0x2c, 0x4a, 0x79, 0x25, 0x06, 0xb3, 0xc7, 0x7e, 0x16, 0x51, 0x8f, 0x59, 0xbb, 0xd1, 0x1e,
	0xdc, 0x33, 0x75, 0xbb, 0x59, 0x0b, 0xc2, 0xb8, 0xe7, 0xcc, 0x3d, 0x67, 0x31, 0x80, 0x8e, 0x64,
	0xaf, 0x4c, 0x53, 0x13, 0x4d, 0x73, 0xce, 0xc4, 0x8f, 0x38, 0xea, 0xae, 0x8c, 0x47, 0xce, 0x2c,
	0x80, 0xde, 0xc3, 0x44, 0xd9, 0x6b, 0xe1, 0x0d, 0x8a, 0x0e, 0x93, 0x3b, 0x1a, 0xb7, 0x4c, 0xec,
	0xf9, 0xe8, 0x13, 0x44, 0xc5, 0xe0, 0x45, 0x63, 0xad, 0x1c, 0x5b, 0xf9, 0xfe, 0xed, 0xf2, 0x6b,
	0xf3, 0xf0, 0xbf, 0x62, 0xeb, 0xe4, 0xa2, 0x20, 0xca, 0xb4, 0xf6, 0x90, 0xee, 0x5d, 0xc4, 0x53,
	0x7f, 0x96, 0x65, 0x5f, 0xac, 0x7f, 0x34, 0xe9, 0xd7, 0xfe, 0xd1, 0x58, 0x79, 0xc7, 0x77, 0x08,
	0x3a, 0x82, 0xa9, 0xbf, 0x21, 0x4d, 0x3c, 0xf3, 0xa7, 0xd8, 0xbb, 0xa3, 0x09, 0xc7, 0xc5, 0x41,
	0xe2, 0xc6, 0xe3, 0x08, 0x79, 0x21, 0x5b, 0x61, 0x62, 0xb0, 0x3b, 0x8f, 0xf1, 0xdc, 0x21, 0x27,
	0x0e, 0x38, 0xfe, 0x35, 0x82, 0xfd, 0x42, 0x6e, 0xfa, 0x8a, 0x15, 0x95, 0x69, 0x5b, 0x15, 0x37,
	0x57, 0x3e, 0x5e, 0x7e, 0xf1, 0xf0, 0x37, 0x0b, 0xaf, 0x3a, 0x14, 0xf7, 0xff, 0x90, 0xef, 0x27,
	0x7d, 0x05, 0xc9, 0x89, 0x7d, 0xdb, 0x52, 0x57, 0x59, 0x45, 0x85, 0x6f, 0x2d, 0xeb, 0x96, 0xec,
	0x9f, 0xa3, 0xb9, 0xf9, 0x57, 0x74, 0x14, 0xc2, 0xf5, 0xd4, 0xf3, 0xdf, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x60, 0x93, 0x02, 0xb9, 0x04, 0x00, 0x00,
}
