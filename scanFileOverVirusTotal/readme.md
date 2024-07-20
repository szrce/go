
example 
```

./main -h
Usage of /var/folders/tj/6ctvsr9d1z5125d90gcf0qk40000gn/T/go-build491164529/b001/exe/main:
  -d string
    	a directory (default "x")
  -k string
    	your api key (default "1")
```


```

output -----
rookie@sezers-MacBook-Pro scanFileOverVirusTotal % go run main.go -d x/ -k apikey
44d88612fea8a8f36de82e1278abb02fFile:  eicar.com.txt
{
    "data": {
        "id": "275a021bbfb6489e54d471899f7db9d1663fc695ec2fe2a2c4538aabf651fd0f",
        "type": "file",
        "links": {
            "self": "https://www.virustotal.com/api/v3/files/275a021bbfb6489e54d471899f7db9d1663fc695ec2fe2a2c4538aabf651fd0f"
        },
        "attributes": {
            "type_tags": [
                "text"
            ],
            "unique_sources": 3557,
            "first_submission_date": 1148301722,
            "type_tag": "text",
            "md5": "44d88612fea8a8f36de82e1278abb02f",
            "times_submitted": 1043553,
            "total_votes": {
                "harmless": 2145,
                "malicious": 383
            },
            "crowdsourced_yara_results": [
                {
                    "ruleset_id": "0019ab4291",
                    "rule_name": "malw_eicar",
                    "ruleset_name": "MALW_Eicar",
                    "description": "Rule to detect the EICAR pattern",
                    "author": "Marc Rivero | McAfee ATR Team",
                    "match_date": 1721250761,
                    "source": "https://github.com/advanced-threat-research/Yara-Rules"
                },
                {
                    "ruleset_id": "015dce072d",
                    "rule_name": "Multi_EICAR_ac8f42d6",
                    "ruleset_name": "Multi_EICAR",
                    "author": "Elastic Security",
                    "match_date": 1721250761,
                    "source": "https://github.com/elastic/protections-artifacts"
                },
                {
                    "ruleset_id": "000720c1f3",
                    "rule_name": "SUSP_Just_EICAR",
                    "ruleset_name": "gen_suspicious_strings",
                    "description": "Just an EICAR test file - this is boring but users asked for it",
                    "author": "Florian Roth (Nextron Systems)",
                    "match_date": 1721250761,
                    "source": "https://github.com/Neo23x0/signature-base"
                }
            ],
            "known_distributors": {
                "distributors": [
                    "Open Source"
                ],
                "filenames": [
                    "eicar.com"
                ],
                "products": [
                    "BlackArch Linux"
                ],
                "data_sources": [
                    "National Software Reference Library (NSRL)"
                ]
            },
            "sigma_analysis_stats": {
                "high": 0,
                "medium": 0,
                "critical": 0,
                "low": 1
            },
            "trid": [
                {
                    "file_type": "EICAR antivirus test file",
                    "probability": 100.0
                }
            ],
            "sha256": "275a021bbfb6489e54d471899f7db9d1663fc695ec2fe2a2c4538aabf651fd0f",
            "ssdeep": "3:a+JraNvsgzsVqSwHq9:tJuOgzsko",
            "last_analysis_results": {
                "Bkav": {
                    "method": "blacklist",
                    "engine_name": "Bkav",
                    "engine_version": "2.0.0.1",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "W32.EicarTest.Trojan"
                },
                "Lionic": {
                    "method": "blacklist",
                    "engine_name": "Lionic",
                    "engine_version": "8.16",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Test.File.EICAR.y"
                },
                "Elastic": {
                    "method": "blacklist",
                    "engine_name": "Elastic",
                    "engine_version": "4.0.157",
                    "engine_update": "20240703",
                    "category": "malicious",
                    "result": "eicar"
                },
                "Cynet": {
                    "method": "blacklist",
                    "engine_name": "Cynet",
                    "engine_version": "4.0.1.1",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Malicious (score: 99)"
                },
                "CMC": {
                    "method": "blacklist",
                    "engine_name": "CMC",
                    "engine_version": "2.4.2022.1",
                    "engine_update": "20240717",
                    "category": "undetected",
                    "result": null
                },
                "CAT-QuickHeal": {
                    "method": "blacklist",
                    "engine_name": "CAT-QuickHeal",
                    "engine_version": "22.00",
                    "engine_update": "20240716",
                    "category": "malicious",
                    "result": "EICAR.TestFile"
                },
                "Skyhigh": {
                    "method": "blacklist",
                    "engine_name": "Skyhigh",
                    "engine_version": "v2021.2.0+4045",
                    "engine_update": "20240716",
                    "category": "malicious",
                    "result": "EICAR test file"
                },
                "ALYac": {
                    "method": "blacklist",
                    "engine_name": "ALYac",
                    "engine_version": "2.0.0.10",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Misc.Eicar-Test-File"
                },
                "Malwarebytes": {
                    "method": "blacklist",
                    "engine_name": "Malwarebytes",
                    "engine_version": "4.5.5.54",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-AV-Test"
                },
                "Zillya": {
                    "method": "blacklist",
                    "engine_name": "Zillya",
                    "engine_version": "2.0.0.5154",
                    "engine_update": "20240715",
                    "category": "malicious",
                    "result": "EICAR.TestFile"
                },
                "Sangfor": {
                    "method": "blacklist",
                    "engine_name": "Sangfor",
                    "engine_version": "2.25.10.0",
                    "engine_update": "20240711",
                    "category": "malicious",
                    "result": "EICAR-Test-File (not a virus)"
                },
                "CrowdStrike": {
                    "method": "blacklist",
                    "engine_name": "CrowdStrike",
                    "engine_version": "1.0",
                    "engine_update": "20231026",
                    "category": "undetected",
                    "result": null
                },
                "Alibaba": {
                    "method": "blacklist",
                    "engine_name": "Alibaba",
                    "engine_version": "0.3.0.5",
                    "engine_update": "20190527",
                    "category": "malicious",
                    "result": "Trojan:MacOS/eicar.com"
                },
                "K7GW": {
                    "method": "blacklist",
                    "engine_name": "K7GW",
                    "engine_version": "12.176.52625",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR_Test_File"
                },
                "K7AntiVirus": {
                    "method": "blacklist",
                    "engine_name": "K7AntiVirus",
                    "engine_version": "12.176.52625",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR_Test_File"
                },
                "Arcabit": {
                    "method": "blacklist",
                    "engine_name": "Arcabit",
                    "engine_version": "2022.0.0.18",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File (not a virus)"
                },
                "Baidu": {
                    "method": "blacklist",
                    "engine_name": "Baidu",
                    "engine_version": "1.0.0.2",
                    "engine_update": "20190318",
                    "category": "malicious",
                    "result": "Win32.Test.Eicar.a"
                },
                "VirIT": {
                    "method": "blacklist",
                    "engine_name": "VirIT",
                    "engine_version": "9.5.748",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File"
                },
                "SymantecMobileInsight": {
                    "method": "blacklist",
                    "engine_name": "SymantecMobileInsight",
                    "engine_version": "2.0",
                    "engine_update": "20240103",
                    "category": "malicious",
                    "result": "ALG:EICAR Test String"
                },
                "Symantec": {
                    "method": "blacklist",
                    "engine_name": "Symantec",
                    "engine_version": "1.21.0.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR Test String"
                },
                "tehtris": {
                    "method": "blacklist",
                    "engine_name": "tehtris",
                    "engine_version": "v0.1.4",
                    "engine_update": "20240717",
                    "category": "undetected",
                    "result": null
                },
                "ESET-NOD32": {
                    "method": "blacklist",
                    "engine_name": "ESET-NOD32",
                    "engine_version": "29573",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Eicar test file"
                },
                "Zoner": {
                    "method": "blacklist",
                    "engine_name": "Zoner",
                    "engine_version": "2.2.2.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR.Test.File-NoVirus.250"
                },
                "APEX": {
                    "method": "blacklist",
                    "engine_name": "APEX",
                    "engine_version": "6.553",
                    "engine_update": "20240716",
                    "category": "malicious",
                    "result": "EICAR Anti-Virus Test File"
                },
                "Avast": {
                    "method": "blacklist",
                    "engine_name": "Avast",
                    "engine_version": "23.9.8494.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR Test-NOT virus!!!"
                },
                "ClamAV": {
                    "method": "blacklist",
                    "engine_name": "ClamAV",
                    "engine_version": "1.3.1.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Win.Test.EICAR_HDB-1"
                },
                "Kaspersky": {
                    "method": "blacklist",
                    "engine_name": "Kaspersky",
                    "engine_version": "22.0.1.28",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File"
                },
                "BitDefender": {
                    "method": "blacklist",
                    "engine_name": "BitDefender",
                    "engine_version": "7.2",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File (not a virus)"
                },
                "NANO-Antivirus": {
                    "method": "blacklist",
                    "engine_name": "NANO-Antivirus",
                    "engine_version": "1.0.146.25796",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Marker.Dos.EICAR-Test-File.dyb"
                },
                "ViRobot": {
                    "method": "blacklist",
                    "engine_name": "ViRobot",
                    "engine_version": "2014.3.20.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-test"
                },
                "MicroWorld-eScan": {
                    "method": "blacklist",
                    "engine_name": "MicroWorld-eScan",
                    "engine_version": "14.0.409.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File"
                },
                "Tencent": {
                    "method": "blacklist",
                    "engine_name": "Tencent",
                    "engine_version": "1.0.0.1",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR.TEST.NOT-A-VIRUS"
                },
                "TACHYON": {
                    "method": "blacklist",
                    "engine_name": "TACHYON",
                    "engine_version": "2024-07-17.02",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File"
                },
                "Emsisoft": {
                    "method": "blacklist",
                    "engine_name": "Emsisoft",
                    "engine_version": "2024.1.0.53752",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File (A)"
                },
                "F-Secure": {
                    "method": "blacklist",
                    "engine_name": "F-Secure",
                    "engine_version": "18.10.1547.307",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR_Test_File"
                },
                "DrWeb": {
                    "method": "blacklist",
                    "engine_name": "DrWeb",
                    "engine_version": "7.0.65.5230",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR Test File (NOT a Virus!)"
                },
                "VIPRE": {
                    "method": "blacklist",
                    "engine_name": "VIPRE",
                    "engine_version": "6.0.0.35",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File (not a virus)"
                },
                "TrendMicro": {
                    "method": "blacklist",
                    "engine_name": "TrendMicro",
                    "engine_version": "11.0.0.1006",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Eicar_test_file"
                },
                "FireEye": {
                    "method": "blacklist",
                    "engine_name": "FireEye",
                    "engine_version": "35.47.0.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File (not a virus)"
                },
                "Sophos": {
                    "method": "blacklist",
                    "engine_name": "Sophos",
                    "engine_version": "2.5.5.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-AV-Test"
                },
                "SentinelOne": {
                    "method": "blacklist",
                    "engine_name": "SentinelOne",
                    "engine_version": "24.2.1.1",
                    "engine_update": "20240417",
                    "category": "malicious",
                    "result": "Static AI - Malicious COM"
                },
                "GData": {
                    "method": "blacklist",
                    "engine_name": "GData",
                    "engine_version": "A:25.38517B:27.36755",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR_TEST_FILE"
                },
                "Jiangmin": {
                    "method": "blacklist",
                    "engine_name": "Jiangmin",
                    "engine_version": "16.0.100",
                    "engine_update": "20240716",
                    "category": "malicious",
                    "result": "EICAR-Test-File"
                },
                "Webroot": {
                    "method": "blacklist",
                    "engine_name": "Webroot",
                    "engine_version": "1.0.0.403",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "W32.Eicar.Testvirus.Gen"
                },
                "Google": {
                    "method": "blacklist",
                    "engine_name": "Google",
                    "engine_version": "1721241039",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Detected"
                },
                "Avira": {
                    "method": "blacklist",
                    "engine_name": "Avira",
                    "engine_version": "8.3.3.20",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Eicar-Test-Signature"
                },
                "Antiy-AVL": {
                    "method": "blacklist",
                    "engine_name": "Antiy-AVL",
                    "engine_version": "3.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "TestFile/Win32.EICAR"
                },
                "Kingsoft": {
                    "method": "blacklist",
                    "engine_name": "Kingsoft",
                    "engine_version": "None",
                    "engine_update": "20230906",
                    "category": "malicious",
                    "result": "Test.eicar.aa"
                },
                "Gridinsoft": {
                    "method": "blacklist",
                    "engine_name": "Gridinsoft",
                    "engine_version": "1.0.182.174",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Trojan.U.EICAR_Test_File.dd"
                },
                "Xcitium": {
                    "method": "blacklist",
                    "engine_name": "Xcitium",
                    "engine_version": "36883",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Malware@#2975xfk8s2pq1"
                },
                "Microsoft": {
                    "method": "blacklist",
                    "engine_name": "Microsoft",
                    "engine_version": "1.1.24060.5",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Virus:DOS/EICAR_Test_File"
                },
                "SUPERAntiSpyware": {
                    "method": "blacklist",
                    "engine_name": "SUPERAntiSpyware",
                    "engine_version": "5.6.0.1032",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "NotAThreat.EICAR[TestFile]"
                },
                "ZoneAlarm": {
                    "method": "blacklist",
                    "engine_name": "ZoneAlarm",
                    "engine_version": "1.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File"
                },
                "Avast-Mobile": {
                    "method": "blacklist",
                    "engine_name": "Avast-Mobile",
                    "engine_version": "240717-00",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Eicar"
                },
                "Varist": {
                    "method": "blacklist",
                    "engine_name": "Varist",
                    "engine_version": "6.5.1.2",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR_Test_File"
                },
                "AhnLab-V3": {
                    "method": "blacklist",
                    "engine_name": "AhnLab-V3",
                    "engine_version": "3.26.0.10499",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Virus/EICAR_Test_File"
                },
                "Acronis": {
                    "method": "blacklist",
                    "engine_name": "Acronis",
                    "engine_version": "1.2.0.121",
                    "engine_update": "20240328",
                    "category": "undetected",
                    "result": null
                },
                "McAfee": {
                    "method": "blacklist",
                    "engine_name": "McAfee",
                    "engine_version": "6.0.6.653",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR test file"
                },
                "MAX": {
                    "method": "blacklist",
                    "engine_name": "MAX",
                    "engine_version": "2023.1.4.1",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "malware (ai score=100)"
                },
                "VBA32": {
                    "method": "blacklist",
                    "engine_name": "VBA32",
                    "engine_version": "5.0.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File"
                },
                "TrendMicro-HouseCall": {
                    "method": "blacklist",
                    "engine_name": "TrendMicro-HouseCall",
                    "engine_version": "10.0.0.1040",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "Eicar_test_file"
                },
                "Rising": {
                    "method": "blacklist",
                    "engine_name": "Rising",
                    "engine_version": "25.0.0.27",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File (CLASSIC)"
                },
                "Yandex": {
                    "method": "blacklist",
                    "engine_name": "Yandex",
                    "engine_version": "5.5.2.24",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR_test_file"
                },
                "Ikarus": {
                    "method": "blacklist",
                    "engine_name": "Ikarus",
                    "engine_version": "6.3.12.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-Test-File"
                },
                "MaxSecure": {
                    "method": "blacklist",
                    "engine_name": "MaxSecure",
                    "engine_version": "1.0.0.1",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "VIRUS.EICAR.TEST"
                },
                "Fortinet": {
                    "method": "blacklist",
                    "engine_name": "Fortinet",
                    "engine_version": "None",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR_TEST_FILE"
                },
                "BitDefenderTheta": {
                    "method": "blacklist",
                    "engine_name": "BitDefenderTheta",
                    "engine_version": "7.2.37796.0",
                    "engine_update": "20240621",
                    "category": "malicious",
                    "result": "EICAR-Test-File (not a virus)"
                },
                "AVG": {
                    "method": "blacklist",
                    "engine_name": "AVG",
                    "engine_version": "23.9.8494.0",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR Test-NOT virus!!!"
                },
                "Cybereason": {
                    "method": "blacklist",
                    "engine_name": "Cybereason",
                    "engine_version": "1.2.449",
                    "engine_update": "20240707",
                    "category": "undetected",
                    "result": null
                },
                "Panda": {
                    "method": "blacklist",
                    "engine_name": "Panda",
                    "engine_version": "4.6.4.2",
                    "engine_update": "20240717",
                    "category": "malicious",
                    "result": "EICAR-AV-TEST-FILE"
                },
                "alibabacloud": {
                    "method": "blacklist",
                    "engine_name": "alibabacloud",
                    "engine_version": "2.1.0",
                    "engine_update": "20240620",
                    "category": "malicious",
                    "result": "Engtest:multi/Eicar test file.Gen"
                },
                "BitDefenderFalx": {
                    "method": "blacklist",
                    "engine_name": "BitDefenderFalx",
                    "engine_version": "2.0.936",
                    "engine_update": "20240128",
                    "category": "type-unsupported",
                    "result": null
                },
                "McAfeeD": {
                    "method": "blacklist",
                    "engine_name": "McAfeeD",
                    "engine_version": "1.2.0.7977",
                    "engine_update": "20240717",
                    "category": "type-unsupported",
                    "result": null
                },
                "DeepInstinct": {
                    "method": "blacklist",
                    "engine_name": "DeepInstinct",
                    "engine_version": "5.0.0.8",
                    "engine_update": "20240715",
                    "category": "type-unsupported",
                    "result": null
                },
                "Trapmine": {
                    "method": "blacklist",
                    "engine_name": "Trapmine",
                    "engine_version": "4.0.16.173",
                    "engine_update": "20240712",
                    "category": "type-unsupported",
                    "result": null
                },
                "Paloalto": {
                    "method": "blacklist",
                    "engine_name": "Paloalto",
                    "engine_version": "0.9.0.1003",
                    "engine_update": "20240717",
                    "category": "type-unsupported",
                    "result": null
                },
                "Cylance": {
                    "method": "blacklist",
                    "engine_name": "Cylance",
                    "engine_version": "3.0.0.0",
                    "engine_update": "20240711",
                    "category": "type-unsupported",
                    "result": null
                },
                "Trustlook": {
                    "method": "blacklist",
                    "engine_name": "Trustlook",
                    "engine_version": "1.0",
                    "engine_update": "20240717",
                    "category": "type-unsupported",
                    "result": null
                }
            },
            "sandbox_verdicts": {
                "Zenbox": {
                    "category": "harmless",
                    "malware_classification": [
                        "CLEAN"
                    ],
                    "sandbox_name": "Zenbox",
                    "confidence": 100
                },
                "Lastline": {
                    "category": "malicious",
                    "sandbox_name": "Lastline",
                    "malware_classification": [
                        "MALWARE",
                        "TROJAN"
                    ]
                },
                "OS X Sandbox": {
                    "category": "malicious",
                    "confidence": 52,
                    "sandbox_name": "OS X Sandbox",
                    "malware_classification": [
                        "MALWARE",
                        "TROJAN",
                        "EVADER"
                    ],
                    "malware_names": [
                        "EICAR"
                    ]
                }
            },
            "magika": "VBA",
            "sigma_analysis_results": [
                {
                    "match_context": [
                        {
                            "values": {
                                "EventID": "4672",
                                "PrivilegeList": "SeAssignPrimaryTokenPrivilege\r\n\t\t\tSeTcbPrivilege\r\n\t\t\tSeSecurityPrivilege\r\n\t\t\tSeTakeOwnershipPrivilege\r\n\t\t\tSeLoadDriverPrivilege\r\n\t\t\tSeBackupPrivilege\r\n\t\t\tSeRestorePrivilege\r\n\t\t\tSeDebugPrivilege\r\n\t\t\tSeAuditPrivilege\r\n\t\t\tSeSystemEnvironmentPrivilege\r\n\t\t\tSeImpersonatePrivilege\r\n\t\t\tSeDelegateSessionUserImpersonatePrivilege",
                                "SubjectUserName": "SYSTEM",
                                "SubjectLogonId": "999",
                                "SubjectUserSid": "S-1-5-18",
                                "SubjectDomainName": "NT AUTHORITY"
                            }
                        },
                        {
                            "values": {
                                "EventID": "4672",
                                "PrivilegeList": "SeAssignPrimaryTokenPrivilege\r\n\t\t\tSeTcbPrivilege\r\n\t\t\tSeSecurityPrivilege\r\n\t\t\tSeTakeOwnershipPrivilege\r\n\t\t\tSeLoadDriverPrivilege\r\n\t\t\tSeBackupPrivilege\r\n\t\t\tSeRestorePrivilege\r\n\t\t\tSeDebugPrivilege\r\n\t\t\tSeAuditPrivilege\r\n\t\t\tSeSystemEnvironmentPrivilege\r\n\t\t\tSeImpersonatePrivilege",
                                "SubjectUserName": "SYSTEM",
                                "SubjectLogonId": "999",
                                "SubjectUserSid": "S-1-5-18",
                                "SubjectDomainName": "NT AUTHORITY"
                            }
                        }
                    ],
                    "rule_level": "low",
                    "rule_description": "Detects logon with \"Special groups\" and \"Special Privileges\" can be thought of as Administrator groups or privileges.",
                    "rule_source": "Sigma Integrated Rule Set (GitHub)",
                    "rule_title": "User with Privileges Logon",
                    "rule_id": "8919a871f4a52b7af785fab44b4665ab6a3637e6ebeeac0288df8a5012a48be2",
                    "rule_author": "frack113"
                }
            ],
            "last_analysis_stats": {
                "malicious": 66,
                "suspicious": 0,
                "undetected": 5,
                "harmless": 0,
                "timeout": 0,
                "confirmed-timeout": 0,
                "failure": 0,
                "type-unsupported": 7
            },
            "type_extension": "txt",
            "meaningful_name": "eicar.com-32189",
            "type_description": "Text",
            "sha1": "3395856ce81f2b7382dee72602f798b642f14140",
            "antiy_info": "Trojan/Generic.ASBOL.2A",
            "magic": "EICAR virus test files",
            "reputation": 3658,
            "last_modification_date": 1721250791,
            "last_submission_date": 1721250760,
            "first_seen_itw_date": 1582585760,
            "tlsh": "T141A022003B0EEE2BA20B00200032E8B00808020E2CE00A3820A020B8C83308803EC228",
            "popular_threat_classification": {
                "popular_threat_category": [
                    {
                        "count": 14,
                        "value": "virus"
                    },
                    {
                        "count": 3,
                        "value": "trojan"
                    }
                ],
                "popular_threat_name": [
                    {
                        "count": 60,
                        "value": "eicar"
                    },
                    {
                        "count": 52,
                        "value": "test"
                    },
                    {
                        "count": 39,
                        "value": "file"
                    }
                ],
                "suggested_threat_label": "virus.eicar/test"
            },
            "last_analysis_date": 1721250760,
            "names": [
                "eicar.com-32189",
                "eicar.com-38688",
                "eicar.com.txt",
                "eicar.com-18159",
                "eicar.com-28964",
                "eicar.com-4151",
                "lasauce1.txt",
                "eicar.com-21000",
                "eicar.txt",
                "eicar.com-29233",
                "eicar.com-12802",
                "rad61800.tmp.com",
                "eicar.com-15256",
                "eicar.com-4817",
                "eicar.com-47533",
                "eicar.com",
                "eicar-test.txt",
                "eicar.com-33247",
                "eicar.com-36231",
                "eicar.com-18850",
                "eicar.com-28284",
                "eicar.com-2880",
                "Confidentail.txt",
                "eicar.com-37041",
                "eicar.com-12150",
                "eicar.com-4156",
                "eicar.com-20218",
                "New Text Document.txt",
                "eicar.com-44930",
                "eicar.com-4522",
                "eicar.com-36986",
                "eicar.com-37017",
                "eicar.com-27644",
                "eicar.com-20443",
                "eicar.com-19393",
                "eicar.com-5963",
                "eicar.com-10115",
                "eicar.com-40491",
                "test.pdf",
                "eicar.com-2114",
                "eicar.com-25749",
                "eicar.pdf",
                "eicar.com-42746",
                "eicar.com-9803",
                "Test_Eicar_P.png",
                "malicous.php",
                "eicar.com-34800",
                "eicar.com-43057",
                "eicar.com-26836",
                "eicar.com-26973",
                "eicar-exe.com",
                "eicar.com-18574",
                "eicar.com-11704",
                "downloadFile.xlsx",
                "eicar.com-10638",
                "eicar.dll",
                "eicar.com-34762",
                "eicar.com-2626",
                "eicar.com-20716",
                "eicar.com-43403",
                "eicar.com-4837",
                "eicar.com-35445",
                "eicar.com-35286",
                "eicar.com-27512",
                "eicar.com-20532",
                "eicar_00.com",
                "eicar_01.com",
                "eicar_02.com",
                "eicar.com-19313",
                "eicar.com-6562",
                "eicar.com-11313",
                "1",
                "eicar.com-40949",
                "eicar.com-3336",
                "eicar.com-27202",
                "eicar.com-44107",
                "eicar.com-13234",
                "eicar.com-36172",
                "eicar.com-42855",
                "eicar.com-28210",
                "eicar.com-28889",
                "eicar.com-20237",
                "eicar.com-14591",
                "eicar.com-12084",
                "eicar.com-48752",
                "eicar.com-31610",
                "eicar.com-16122",
                "eicar.com-1559",
                "eicar.com-28931",
                "eicar.com-35726",
                "eicar.com-20928",
                "eicar.com-31050",
                "eicar.com-26490",
                "eicar.com-15444",
                "eicar.com-12476",
                "eicar.com-7470",
                "eicar.com-46810"
            ],
            "sigma_analysis_summary": {
                "Sigma Integrated Rule Set (GitHub)": {
                    "high": 0,
                    "medium": 0,
                    "critical": 0,
                    "low": 1
                }
            },
            "tags": [
                "direct-cpu-clock-access",
                "idle",
                "legit",
                "known-distributor",
                "attachment",
                "via-tor",
                "long-sleeps"
            ],
            "size": 68,
            "crowdsourced_ai_results": [
                {
                    "analysis": "EICAR is a test string used to detect and test antivirus software. It's like a \"dummy virus\" that triggers an antivirus engine to react, demonstrating its ability to identify and neutralize threats.\nHere's the key:\nIt's NOT a real virus: EICAR is harmless and cannot infect your computer.\nIt's a standardized test: Almost all antivirus programs are designed to recognize EICAR as a potential threat, ensuring they're working properly.\n",
                    "source": "palm",
                    "category": "code_insight",
                    "id": "275a021bbfb6489e54d471899f7db9d1663fc695ec2fe2a2c4538aabf651fd0f-file-palm"
                }
            ]
        }
    }
}
sezer@sezers-MacBook-Pro scanFileOverVirusTotal % 
