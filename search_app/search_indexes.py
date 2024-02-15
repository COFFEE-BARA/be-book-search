# search_app/search_indexes.py

from django_elasticsearch_dsl import Document, Index
from elasticsearch_dsl import analyzer, tokenizer
from .models import Book

# Elasticsearch DSL을 사용하여 인덱스 정의
book_ngram = Index('book_ngram')

# 분석기 선택 함수
def combined_analyzer(language):
    if language == 'english':
        return analyzer(analyzer=combined_analyzer, tokenizer=tokenizer('ngram', 'nGram', min_gram=2, max_gram=20), filter=['lowercase'])
    elif language == 'korean':
        return analyzer('korean_analyzer', tokenizer='nori_tokenizer', filter=['ngram_filter'])

# 모델과 연결
@book_ngram.doc_type
class BookDocument(Document):
    class Index:
        name = 'book_ngram'
        settings = {
            'number_of_shards': 1,
            'number_of_replicas': 0,
            'max_ngram_diff': 20,
            'analysis': {
                'tokenizer': {
                    'ngram_tokenizer': {
                        'type': 'ngram',
                        'min_gram': 1,
                        'max_gram': 20
                    }
                },
                'analyzer': {
                    'combined_analyzer': combined_analyzer,
                    'korean_analyzer': {
                        'type': 'custom',
                        'tokenizer': 'nori_tokenizer',
                        'filter': ['ngram_filter']
                    }
                },
                'filter': {
                    'ngram_filter': {
                        'type': 'ngram',
                        'min_gram': 1,
                        'max_gram': 20
                    }
                }
            }
        }

    class Django:
        model = Book
        fields = [
            'Title'
        ]

    def get_analyzer(self, **kwargs):
        # 언어에 따라 분석기 선택
        return combined_analyzer(kwargs.get('language', 'english'))
