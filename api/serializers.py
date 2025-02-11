from rest_framework import serializers
from django.contrib.auth.models import User
from .models import Recipe

class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ['id', 'username']

class RecipeSerializer(serializers.ModelSerializer):
    author = UserSerializer()

    class Meta:
        model = Recipe
        fields = ['id', 'title', 'ingredients', 'instructions', 'author']
        depth = 1
