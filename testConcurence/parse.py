import re
from resource_map import RESOURCE_NAME_MAP
import mysql.connector

filename = 'pipeline_search_results_20230216_112743.txt'
db = mysql.connector.connect(
  host = "localhost",
  user = "root",
  database = "playground_db",
  use_pure = True,
)

with open(filename, 'rb') as f:
    line_num = 0
    text_payload_idx = 0

    for line in f:
        line_num += 1
        if line_num == 1:
            continue

        line = line.strip().replace('-\t-\t-\t-\t', '')
        split_line = line.split('\t')

        if line_num == 2:
            for i in range(len(split_line)):
                if split_line[i] == '@text_payload':
                    text_payload_idx = i
                    break
            continue

        content = split_line[text_payload_idx].split('|[3pl monitor]')

        ctx_content = content[0]
        ctx_request_id = ctx_content.split('|')[-1]

        log_content = content[1]

        # pattern for recordCreateAndCancelLog
        pattern = 'requestId=(.*),' + \
            'region=(?P<region>.*),' + \
            'clientSn=(?P<clientSn>.*),' + \
            'tplTn=(?P<tplTn>.*),' + \
            'clientId=(?P<clientId>.*),' + \
            'resourceId=(?P<resourceId>.*),' + \
            'lineId=(?P<lineId>.*),' + \
            'realCall3PLFlag=(?P<realCall3PLFlag>.*),' + \
            'opType=(?P<opType>.*),' + \
            'elapse=(?P<elapse>.*),' + \
            'ts=(?P<ts>.*),' + \
            'protoType=(?P<protoType>.*),' + \
            'httpStatus=(?P<httpStatus>.*),' + \
            'url=(?P<url>.*),' + \
            'startTs=(?P<startTs>.*),' + \
            'endTs=(?P<endTs>.*),' + \
            'body=(?P<body>.*),' + \
            'retCode=(?P<retCode>.*),' + \
            'message=(?P<message>.*),' + \
            'type=(?P<type>.*),' + \
            'apiCode=(?P<apiCode>.*),' + \
            'apiMessage=(?P<apiMessage>.*),' + \
            'reply=(?P<reply>.*)'

        log_dict = re.search(pattern, log_content).groupdict()

        cur = db.cursor()
        sql = 'INSERT INTO pis_3pl_error_log_tab (' + \
            'request_id,' + \
            'region,' + \
            'client_sn,' + \
            'resource_id,' + \
            'resource_name,' + \
            'op_type,' + \
            'url,' + \
            'http_status,' + \
            'retcode,' + \
            'message,' + \
            'api_code,' + \
            'api_message,' + \
            'reply,' + \
            'timestamp) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)' + \
            'ON DUPLICATE KEY UPDATE request_id=request_id, client_sn=client_sn'
        row = (
            ctx_request_id,
            log_dict['region'],
            log_dict['clientSn'],
            log_dict['resourceId'],
            RESOURCE_NAME_MAP[log_dict['resourceId']],
            log_dict['opType'],
            log_dict['url'],
            log_dict['httpStatus'],
            log_dict['retCode'],
            log_dict['message'],
            log_dict['apiCode'],
            log_dict['apiMessage'],
            log_dict['reply'],
            log_dict['ts']
        )

        print(filename, ctx_request_id, log_dict['clientSn'])

        cur.execute(sql, row)
        db.commit()

db.close()