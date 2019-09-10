# removeStop

## Overview

### Description
Given text and list of (stop) words - removes those words from text.

### Implementation details
Uses the lib: github.com/bbalet/stopwords.  To include punctuation the regex used as a word segmenter was over writen to keep punctuation.

## Compliance to Spec

### Rough level of compliance  

90%

### Deviations from Specification
- Write now the stopwords from the library do not match the stop words from NLTK that I used as a baseline