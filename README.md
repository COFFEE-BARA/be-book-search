# 📚 Checkbara
![checkbara1](https://github.com/COFFEE-BARA/be-library-stock/assets/72396865/5a3da8b6-009d-40f5-a3f8-a1561df83631)
![checkbara2](https://github.com/COFFEE-BARA/be-library-stock/assets/72396865/eeef4f87-fb54-471b-bbba-2f9fac767823)
![checkbara3](https://github.com/COFFEE-BARA/be-library-stock/assets/72396865/3a8da5c0-19ee-4251-9161-96a835ce0b3f)
<img width="875" alt="checkbara4" src="https://github.com/COFFEE-BARA/be-library-stock/assets/72396865/fe97f585-7be6-4810-bc0a-f3c331ce9c39">

<br/>

| <img width="165" alt="suwha" src="https://github.com/COFFEE-BARA/be-bookstore-stock/assets/72396865/19e01fac-5384-4ec7-98f1-9e1e613429b4"> | <img width="165" alt="yoonju" src="https://github.com/COFFEE-BARA/be-bookstore-stock/assets/72396865/fb0a14c6-2d02-4105-962e-4565663817cc"> | <img width="165" alt="yugyeong" src="https://github.com/COFFEE-BARA/be-bookstore-stock/assets/72396865/90b7268d-92e5-43d1-9da8-ae48afd9e8c1"> | <img width="165" alt="dayeon" src="https://github.com/COFFEE-BARA/be-bookstore-stock/assets/72396865/f19e65e6-0856-4b6a-a355-993ce83ddcb7"> |
| --- | --- | --- | --- |
| 🐼[유수화](https://github.com/YuSuhwa-ve)🐼 | 🐱[송윤주](https://github.com/raminicano)🐱 | 🐶[현유경](https://github.com/yugyeongh)🐶 | 🐤[양다연](https://github.com/dayeon1201)🐤 |
| Server / Data / BE | AI / Data / BE | Infra / BE / FE | BE / FE |

<br/>


# 🔎 be-book-search
![__-ezgif com-resize](https://github.com/COFFEE-BARA/crawler-kyobo-isbn/assets/65851554/699d09a6-8691-4ca0-bab5-575340a3c34d)

본 레포의 목표는 사용자가 더욱 정확하고 빠르게 원하는 책을 찾을 수 있도록 검색 기능을 고도화하는 것입니다. 한국어 및 영어 제목 검색과 저자 검색의 정확도와 성능을 향상시키기 위한 여러 전략을 구현했습니다.
<br/>
<br>
<br>

## 기능 소개
### 📕 제목 검색의 고도화

- **한국어 검색 최적화**: `nori_tokenizer`와 필터(`nori_readingform`, `nori_part_of_speech`)를 적용하여 한국어 복합어를 효과적으로 처리하고 검색 정확도를 높였습니다.
- **영어 제목 검색 강화**: 모든 영어 제목에 대한 대소문자 구분 없이 검색할 수 있도록 `lowercase_filter`를 적용했습니다.
- **커스텀 분석기 적용**: 이러한 설정을 포함한 커스텀 분석기를 제목 필드에 적용하여, 제목 검색 기능을 고도화했습니다.

### 📙 저자 검색의 고도화

- **부분 일치 검색 가능**: `edge_ngram_filter`를 적용하여 저자 이름의 부분 일치 검색을 가능하게 했습니다. 사용자는 쿼리의 일부만으로도 저자를 검색할 수 있습니다.
- **대소문자 구분 없는 검색**: 저자 검색 시 대소문자를 구분하지 않도록 `lowercase_filter`를 적용했습니다.

### 📗 검색 시스템의 성능 최적화

- **백엔드 성능 향상**: 검색 시스템의 백엔드를 Go 언어로 작성하여 프론트엔드와의 통신 시간을 대폭 줄였습니다. Go의 빠른 실행 속도 덕분에 기존 시스템 대비 통신 시간을 최대 5배까지 단축했습니다.
- **빠른 반응 시간**: 이로 인해 사용자는 더 빠른 반응 시간으로 검색 결과를 받아볼 수 있습니다.
- **복합 쿼리 검색**: `multi_match` 쿼리를 활용하여 사용자가 제목과 저자를 동시에 검색할 수 있는 기능을 구현했습니다.

결과적으로 이러한 고도화 작업을 통해 사용자 경험을 대폭 향상시키고, 원하는 책을 더욱 빠르고 정확하게 찾을 수 있는 검색 시스템을 제공합니다.


<br/>
<br/>


## 📝 Elastic index
<details>
<summary>book-index mapping</summary>
<div markdown="1">

```
// book-index mapping

{
  "mappings": {
    "properties": {
      "Author": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "text",
            "analyzer": "author_analyzer"
          },
          "partial": {
            "type": "text",
            "analyzer": "edge_ngram_analyzer"
          }
        }
      },
      "DetailCategory": {
        "type": "keyword"
      },
      "ISBN": {
        "type": "keyword"
      },
      "ImageURL": {
        "type": "keyword"
      },
      "IndexContent": {
        "type": "text"
      },
      "Introduction": {
        "type": "text"
      },
      "MiddleCategory": {
        "type": "keyword"
      },
      "Price": {
        "type": "integer"
      },
      "PubDate": {
        "type": "date",
        "format": "yyyy-MM-dd"
      },
      "Publisher": {
        "type": "keyword"
      },
      "PublisherReview": {
        "type": "text"
      },
      "PurchaseURL": {
        "type": "keyword"
      },
      "Search": {
        "type": "text"
      },
      "Title": {
        "type": "text",
        "analyzer": "title_analyzer"
      },
      "Vector": {
        "type": "dense_vector",
        "dims": 768,
        "index": true,
        "similarity": "cosine"
      },
      "document": {
        "type": "object"
      },
      "id": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "index": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "pipeline": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      }
    }
  }
}

```

</div>
</details>

<details>
<summary> book-index settings </summary>
<div markdown="1">

```
//book -index settings
{
  "settings": {
    "index": {
      "routing": {
        "allocation": {
          "include": {
            "_tier_preference": "data_content"
          }
        }
      },
      "number_of_shards": "1",
      "provided_name": "book-index",
      "creation_date": "1708182319595",
      "analysis": {
        "filter": {
          "lowercase_filter": {
            "type": "lowercase"
          },
          "edge_ngram_filter": {
            "type": "edge_ngram",
            "min_gram": "1",
            "max_gram": "10"
          }
        },
        "analyzer": {
          "edge_ngram_analyzer": {
            "filter": [
              "edge_ngram_filter",
              "lowercase_filter"
            ],
            "type": "custom",
            "tokenizer": "nori_tokenizer_mine"
          },
          "author_analyzer": {
            "filter": [
              "lowercase_filter"
            ],
            "type": "custom",
            "tokenizer": "keyword"
          },
          "title_analyzer": {
            "filter": [
              "nori_readingform",
              "lowercase_filter",
              "nori_part_of_speech"
            ],
            "type": "custom",
            "tokenizer": "nori_tokenizer_mine"
          }
        },
        "tokenizer": {
          "nori_tokenizer_mine": {
            "type": "nori_tokenizer",
            "decompound_mode": "mixed"
          }
        }
      },
      "number_of_replicas": "2",
      "uuid": "okUbOg_pTJKVG2WO7e3rYQ",
      "version": {
        "created": "8500003"
      }
    }
  },
  "defaults": {
    "index": {
      "flush_after_merge": "512mb",
      "time_series": {
        "end_time": "9999-12-31T23:59:59.999Z",
        "start_time": "-9999-01-01T00:00:00Z",
        "es87tsdb_codec": {
          "enabled": "true"
        }
      },
      "final_pipeline": "_none",
      "max_inner_result_window": "100",
      "unassigned": {
        "node_left": {
          "delayed_timeout": "1m"
        }
      },
      "max_terms_count": "65536",
      "rollup": {
        "source": {
          "name": "",
          "uuid": ""
        }
      },
      "lifecycle": {
        "prefer_ilm": "true",
        "rollover_alias": "",
        "origination_date": "-1",
        "name": "",
        "parse_origination_date": "false",
        "step": {
          "wait_time_threshold": "12h"
        },
        "indexing_complete": "false"
      },
      "mode": "standard",
      "routing_partition_size": "1",
      "force_memory_term_dictionary": "false",
      "max_docvalue_fields_search": "100",
      "merge": {
        "scheduler": {
          "max_thread_count": "1",
          "auto_throttle": "true",
          "max_merge_count": "6"
        },
        "policy": {
          "merge_factor": "32",
          "floor_segment": "2mb",
          "max_merge_at_once_explicit": "30",
          "max_merge_at_once": "10",
          "max_merged_segment": "0b",
          "expunge_deletes_allowed": "10.0",
          "segments_per_tier": "10.0",
          "type": "UNSET",
          "deletes_pct_allowed": "20.0"
        }
      },
      "max_refresh_listeners": "1000",
      "max_regex_length": "1000",
      "load_fixed_bitset_filters_eagerly": "true",
      "number_of_routing_shards": "1",
      "write": {
        "wait_for_active_shards": "1"
      },
      "verified_before_close": "false",
      "mapping": {
        "coerce": "false",
        "nested_fields": {
          "limit": "50"
        },
        "depth": {
          "limit": "20"
        },
        "field_name_length": {
          "limit": "9223372036854775807"
        },
        "total_fields": {
          "limit": "1000"
        },
        "nested_objects": {
          "limit": "10000"
        },
        "ignore_malformed": "false",
        "dimension_fields": {
          "limit": "21"
        }
      },
      "source_only": "false",
      "soft_deletes": {
        "enabled": "true",
        "retention": {
          "operations": "0"
        },
        "retention_lease": {
          "period": "12h"
        }
      },
      "max_script_fields": "32",
      "query": {
        "default_field": [
          "*"
        ],
        "parse": {
          "allow_unmapped_fields": "true"
        }
      },
      "format": "0",
      "frozen": "false",
      "sort": {
        "missing": [],
        "mode": [],
        "field": [],
        "order": []
      },
      "priority": "1",
      "routing_path": [],
      "version": {
        "compatibility": "8500003"
      },
      "codec": "default",
      "max_rescore_window": "10000",
      "bloom_filter_for_id_field": {
        "enabled": "true"
      },
      "max_adjacency_matrix_filters": "100",
      "analyze": {
        "max_token_count": "10000"
      },
      "gc_deletes": "60s",
      "top_metrics_max_size": "10",
      "optimize_auto_generated_id": "true",
      "max_ngram_diff": "1",
      "hidden": "false",
      "translog": {
        "flush_threshold_age": "1m",
        "generation_threshold_size": "64mb",
        "flush_threshold_size": "10gb",
        "sync_interval": "5s",
        "retention": {
          "size": "-1",
          "age": "-1"
        },
        "durability": "REQUEST"
      },
      "auto_expand_replicas": "false",
      "fast_refresh": "false",
      "recovery": {
        "type": ""
      },
      "requests": {
        "cache": {
          "enable": "true"
        }
      },
      "data_path": "",
      "highlight": {
        "max_analyzed_offset": "1000000",
        "weight_matches_mode": {
          "enabled": "true"
        }
      },
      "look_back_time": "2h",
      "routing": {
        "rebalance": {
          "enable": "all"
        },
        "allocation": {
          "disk": {
            "watermark": {
              "ignore": "false"
            }
          },
          "enable": "all",
          "total_shards_per_node": "-1"
        }
      },
      "search": {
        "slowlog": {
          "level": "TRACE",
          "threshold": {
            "fetch": {
              "warn": "-1",
              "trace": "-1",
              "debug": "-1",
              "info": "-1"
            },
            "query": {
              "warn": "-1",
              "trace": "-1",
              "debug": "-1",
              "info": "-1"
            }
          }
        },
        "idle": {
          "after": "30s"
        },
        "throttled": "false"
      },
      "fielddata": {
        "cache": "node"
      },
      "look_ahead_time": "2h",
      "default_pipeline": "_none",
      "max_slices_per_scroll": "1024",
      "shard": {
        "check_on_startup": "false"
      },
      "xpack": {
        "watcher": {
          "template": {
            "version": ""
          }
        },
        "version": "",
        "ccr": {
          "following_index": "false"
        }
      },
      "percolator": {
        "map_unmapped_fields_as_text": "false"
      },
      "allocation": {
        "max_retries": "5",
        "existing_shards_allocator": "gateway_allocator"
      },
      "refresh_interval": "1s",
      "indexing": {
        "slowlog": {
          "reformat": "true",
          "threshold": {
            "index": {
              "warn": "-1",
              "trace": "-1",
              "debug": "-1",
              "info": "-1"
            }
          },
          "source": "1000",
          "level": "TRACE"
        }
      },
      "compound_format": "1gb",
      "blocks": {
        "metadata": "false",
        "read": "false",
        "read_only_allow_delete": "false",
        "read_only": "false",
        "write": "false"
      },
      "max_result_window": "10000",
      "store": {
        "stats_refresh_interval": "10s",
        "type": "",
        "fs": {
          "fs_lock": "native"
        },
        "preload": [],
        "snapshot": {
          "snapshot_name": "",
          "index_uuid": "",
          "cache": {
            "prewarm": {
              "enabled": "true"
            },
            "enabled": "true",
            "excluded_file_types": []
          },
          "repository_uuid": "",
          "uncached_chunk_size": "-1b",
          "delete_searchable_snapshot": "false",
          "index_name": "",
          "partial": "false",
          "blob_cache": {
            "metadata_files": {
              "max_length": "64kb"
            }
          },
          "repository_name": "",
          "snapshot_uuid": ""
        }
      },
      "queries": {
        "cache": {
          "enabled": "true"
        }
      },
      "shard_limit": {
        "group": "normal"
      },
      "warmer": {
        "enabled": "true"
      },
      "downsample": {
        "origin": {
          "name": "",
          "uuid": ""
        },
        "source": {
          "name": "",
          "uuid": ""
        },
        "status": "unknown"
      },
      "override_write_load_forecast": "0.0",
      "max_shingle_diff": "3",
      "query_string": {
        "lenient": "false"
      }
    }
  }
}

```

</div>
</details>


<br/>
<br/>


## 🌎 API 명세

<details><summary> 1️⃣ [GET] BASE_URL/api/book/search?keyword={검색어}
</summary>

: 유저가 검색어를 입력하면 스코어 내림차순으로 책 리스트를 반환한다.


</details>

<details><summary> 2️⃣ Request </summary>

- **☝🏻Request Header**
    
    ```
    Content-Type: application/json
    
    ```
    
- **✌🏻Request Query**
    
    
    | Name | Type | Description | Required |
    | --- | --- | --- | --- |
    | keyword={검색어} | String | 검색할 책 | Required |
    

    
</details>

<details><summary> 3️⃣ Response </summary>

- **☝🏻성공**
    
    
    | Name | Type | Description |
    | --- | --- | --- |
    | code | Integer | 상태 코드 |
    | message | String | 상태 메시지 |
    | data | List | 상태 코드 |
    | - recommendedBookList | List | ai가 추천하는 책, list는 추천순으로 정렬됨. |
    | -- isbn | Long | 책의 13자리 isbn |
    | -- title | String | 책이름 |
    | -- author | String | 책 저자(들) |
    | -- image | String | 책 표지 이미지 url |
    | -- price | Long | 책의 가격(재고검색을 위해 필요) |
    

  ```
   {
    "code": 200,
    "message": "검색 결과를 가져오는데 성공했습니다.",
		"data": {
			 "bookList" :[
								{
										"isbn" : 9791140708116,
										"title" : "아는 만큼 보이는 백엔드 개발 (한 권으로 보는 백엔드 로드맵과 커리어 가이드)",
										"author" : ["정우현", "이인", "김보인"],
										"image" : "https://shopping-phinf.pstatic.net/main_4519670/45196700648.20240114070834.jpg",
										"price" : 490000,
								},
								{
										"isbn" : 9791140708116,
										"title" : "아는 만큼 보이는 백엔드 개발 (한 권으로 보는 백엔드 로드맵과 커리어 가이드)",
										"author" : ["정우현", "이인", "김보인"],
										"image" : "https://shopping-phinf.pstatic.net/main_4519670/45196700648.20240114070834.jpg",
										"price" : 490000,
								
								},
								{
										"isbn" : 9791140708116,
										"title" : "아는 만큼 보이는 백엔드 개발 (한 권으로 보는 백엔드 로드맵과 커리어 가이드)",
										"author" : ["정우현", "이인", "김보인"],
										"image" : "https://shopping-phinf.pstatic.net/main_4519670/45196700648.20240114070834.jpg",
										"price" : 490000,
								},
								
					   ] 
		    }
     }  

  ```
    
- **✌🏻실패**
    - 필요한 값이 없는 경우
        
        ```
        {
            "code": 400,
            "message": "keyword가 없습니다.",
            "data": null
        }
        
        ```
        
    - 서버에러
        
        ```
        {
            "code": 500,
            "message": "서버 에러",
            "data": null
        }
        
        ```
        

</details>

<br/>
<br/>

## 🛠️ 사용 기술 스택
### Programming language & Library

<img src="https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white"/> ![ElasticSearch](https://img.shields.io/badge/-ElasticSearch-005571?style=for-the-badge&logo=elasticsearch)


### DB

<img src="https://img.shields.io/badge/elastic-005571?style=for-the-badge&logo=elastic&logoColor=white">


### Deploy & CI/CD

<img src="https://img.shields.io/badge/lambda-FF9900?style=for-the-badge&logo=awslambda&logoColor=white"/>  <img src="https://img.shields.io/badge/docker-2496ED?style=for-the-badge&logo=docker&logoColor=white"> <img src="https://img.shields.io/badge/ecr-FC4C02?style=for-the-badge&logo=ecr&logoColor=white"> <img src="https://img.shields.io/badge/codebuild-68A51C?style=for-the-badge&logo=codebuild&logoColor=white"> <img src="https://img.shields.io/badge/codepipeline-527FFF?style=for-the-badge&logo=codepipeline&logoColor=white">

### Version Control System

<img src="https://img.shields.io/badge/github-181717?style=for-the-badge&logo=github&logoColor=white"> <img src="https://img.shields.io/badge/git-F05032?style=for-the-badge&logo=git&logoColor=white">

### Communication Tool

<img src="https://img.shields.io/badge/slack-4A154B?style=for-the-badge&logo=slack&logoColor=white"> <img src="https://img.shields.io/badge/notion-000000?style=for-the-badge&logo=notion&logoColor=white">

<br/>
<br/>

