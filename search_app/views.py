# search_app/views.py

from django.views.decorators.csrf import csrf_exempt
from django.utils.decorators import method_decorator

from rest_framework.views import APIView  
from rest_framework.response import Response  
from rest_framework import status  

from rest_framework.authentication import BasicAuthentication, SessionAuthentication
from rest_framework.permissions import IsAuthenticated

from elasticsearch import Elasticsearch  
from django.conf import settings

from .search_indexes import *  # Elasticsearch DSL이 정의된 인덱스 import

from django.http import JsonResponse
from elasticsearch_dsl import Search


elasticsearch_settings = getattr(settings, 'ELASTICSEARCH', {})

@method_decorator(csrf_exempt,name='dispatch')
class SearchView(APIView):
    authentication_classes = [BasicAuthentication, SessionAuthentication]
    permission_classes = [IsAuthenticated]

    def get(self, request):
        es = Elasticsearch(
            cloud_id= '',
            api_key='',
            request_timeout=600
        )

        # 검색어
        search_word = request.query_params.get('search')

        if not search_word:
            return Response(status=status.HTTP_400_BAD_REQUEST, data={'message': 'search word param is missing'})

        docs = es.search(index='book_ngram',
                         body={
                             "query": {
                                 "multi_match": {
                                     "query": search_word,
                                     "analyzer": "combined_analyzer",
                                     "fields": ["Title"]
                                 }
                             }
                         })

        data_list = docs['hits']

        return Response(data_list)
