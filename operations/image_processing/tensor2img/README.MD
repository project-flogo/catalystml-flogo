# img2tensor

## Overview

### Description
Converts a image in mathematical tensor (an array of any number of dimentions with value of integers ranging from 0 to 255) format to an image type (i.e. jpg, png, gif).

### Implementation details
Converts a image (with and without alpha) in rank 3 or 4 tensor of integers to image.image type. With extension and toFileFolder being set images will be saved as file (jpg, png or gif format). The standard "image, image/jpg, image/png, image/gif" golang library is used.

## Compliance to Spec

### Rough level of compliance  

100%