package mi

import "encoding/xml"

func Marshal(envelope QlmEnvelope) ([]byte, error) {
	//objects.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	//objects.NoNamespaceSchemaLocation = "QLMdf.xsd"
	root := struct {
		QlmEnvelope
		XMLName struct{} `xml:"qlmEnvelope"`
	}{QlmEnvelope: envelope}
	return xml.MarshalIndent(root, "", "    ")
}
