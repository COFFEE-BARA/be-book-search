#search_app/models.py

from django.db import models

class Book(models.Model):
    Title = models.CharField(max_length=255, blank=False)

    def __str__(self):
        return self.Title
