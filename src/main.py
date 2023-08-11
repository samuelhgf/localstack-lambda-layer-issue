import time

import pdfkit


def handler(event, context):

    try:
        nfs_link = event.get("url")

        nfs_number = event.get("number")

        file_name = f'{nfs_number}.pdf'

        path = '/tmp/' + file_name

        PDFKIT_CONFIG = pdfkit.configuration(wkhtmltopdf='/opt/bin/wkhtmltopdf')

        pdfkit.from_url(nfs_link, path, configuration=PDFKIT_CONFIG)

        time.sleep(60)

    except Exception as e:
        print(e)
        raise e
