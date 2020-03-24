module github.com/project-flogo/catalystml-flogo/operations/image_processing

require (
	github.com/disintegration/imaging v1.6.1
	github.com/project-flogo/catalystml-flogo/action v0.0.0-20190919183915-24819a1a9a2b
	github.com/project-flogo/core v0.9.3
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.uber.org/multierr v1.2.0 // indirect
	go.uber.org/zap v1.11.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
	github.com/project-flogo/catalystml-flogo/operations/image_processing/grayscale v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/image_processing/img2tensor v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/image_processing/resize v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/image_processing/subsectiontoimage v0.0.0-20191221100507-49a2889fd614
	github.com/project-flogo/catalystml-flogo/operations/image_processing/tensor2image v0.0.0-20191221100507-49a2889fd614
	
)

replace github.com/project-flogo/catalystml-flogo/operations/image_processing/grayscale => ../operations/image_processing/subsectiontoimage
replace github.com/project-flogo/catalystml-flogo/operations/image_processing/tensor2image => ../operations/image_processing/tensor2image
replace github.com/project-flogo/catalystml-flogo/operations/image_processing/resize => ../operations/image_processing/resize
replace github.com/project-flogo/catalystml-flogo/operations/image_processing/img2tensor => ../operations/image_processing/img2tensor
replace github.com/project-flogo/catalystml-flogo/operations/image_processing/grayscale => ../operations/image_processing/grayscale

