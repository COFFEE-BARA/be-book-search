# search_app/urls.py

from search_app import views  
from django.urls import path  
from rest_framework_simplejwt import views as jwt_views
  
urlpatterns = [  
    path('', views.SearchView.as_view()),  
    path('api/token/', jwt_views.TokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('api/token/refresh/', jwt_views.TokenRefreshView.as_view(), name='token_refresh') 
    # token... 왜 있어야 함? 
]