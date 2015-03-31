#-*- coding: utf-8 -*-

import base64

postDataTmpl = """<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http:/www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
<soap12:Body>
    <UploadFile xmlns="http://tempuri.org/">
        <entNo>PTE51001407270000001</entNo>
        <filename>test.xml</filename>
        <content>{0}</content>
    </UploadFile>
</soap12:Body>
</soap12:Envelope>"""

with open("message.xml","rb") as f:
    body = f.read()
encoded = base64.standard_b64encode(body)
postData = postDataTmpl.format(encoded.decode())

with open("postData.xml","w") as f:
    f.write(postData)

