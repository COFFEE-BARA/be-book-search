# 기본 이미지 설정
FROM python:3.11.1

# 작업 디렉토리 설정
WORKDIR /search_app

# 소스 코드를 현재 디렉토리로 복사
COPY . /search_app

# requirements.txt를 사용하여 필요한 의존성 설치
RUN pip install --no-cache-dir -r requirements.txt

# 컨테이너 내에서 사용할 포트 열기 (Django 애플리케이션의 기본 포트)
EXPOSE 8000

# Django 애플리케이션 실행
CMD ["python3", "manage.py", "runserver", "0.0.0.0:9200"]
