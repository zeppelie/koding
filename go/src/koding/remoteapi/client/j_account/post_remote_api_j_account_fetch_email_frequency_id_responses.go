package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJAccountFetchEmailFrequencyIDReader is a Reader for the PostRemoteAPIJAccountFetchEmailFrequencyID structure.
type PostRemoteAPIJAccountFetchEmailFrequencyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountFetchEmailFrequencyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountFetchEmailFrequencyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountFetchEmailFrequencyIDOK creates a PostRemoteAPIJAccountFetchEmailFrequencyIDOK with default headers values
func NewPostRemoteAPIJAccountFetchEmailFrequencyIDOK() *PostRemoteAPIJAccountFetchEmailFrequencyIDOK {
	return &PostRemoteAPIJAccountFetchEmailFrequencyIDOK{}
}

/*PostRemoteAPIJAccountFetchEmailFrequencyIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountFetchEmailFrequencyIDOK struct {
	Payload PostRemoteAPIJAccountFetchEmailFrequencyIDOKBody
}

func (o *PostRemoteAPIJAccountFetchEmailFrequencyIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.fetchEmailFrequency/{id}][%d] postRemoteApiJAccountFetchEmailFrequencyIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountFetchEmailFrequencyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJAccountFetchEmailFrequencyIDOKBody post remote API j account fetch email frequency ID o k body
swagger:model PostRemoteAPIJAccountFetchEmailFrequencyIDOKBody
*/
type PostRemoteAPIJAccountFetchEmailFrequencyIDOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJAccountFetchEmailFrequencyIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO0

	var postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJAccountFetchEmailFrequencyIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO0)

	postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountFetchEmailFrequencyIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j account fetch email frequency ID o k body
func (o *PostRemoteAPIJAccountFetchEmailFrequencyIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
