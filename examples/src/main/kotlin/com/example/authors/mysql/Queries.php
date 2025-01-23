<?php
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

namespace com\example\authors\mysql;

interface Queries {
  public function createAuthor(string $name, ?string $bio): int;
  
  public function deleteAuthor(int $id): void;
  
  public function getAuthor(int $id): ?Author;
  
  /**
  *  @return array<Author>
  */
  public function listAuthors(): array;
  
}
