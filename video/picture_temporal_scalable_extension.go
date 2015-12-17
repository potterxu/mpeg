package video

import "github.com/32bitkid/mpeg/util"

type PictureTemporalScalableExtension struct {
	reference_select_code       uint32 // 2 uimsbf
	forward_temporal_reference  uint32 // 10 uimsbf
	backward_temporal_reference uint32 // 10 uimsbf
}

func picture_temporal_scalable_extension(br util.BitReader32) (*PictureTemporalScalableExtension, error) {
	err := extension_code_check(br, PictureTemporalScalableExtensionID)
	if err != nil {
		return nil, err
	}

	ptse := PictureTemporalScalableExtension{}

	ptse.reference_select_code, err = br.Read32(2)
	if err != nil {
		return nil, err
	}

	ptse.forward_temporal_reference, err = br.Read32(10)
	if err != nil {
		return nil, err
	}

	err = marker_bit(br)
	if err != nil {
		return nil, err
	}

	ptse.backward_temporal_reference, err = br.Read32(10)
	if err != nil {
		return nil, err
	}

	return &ptse, next_start_code(br)

}
