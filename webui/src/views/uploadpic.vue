<template>
    <div class="image-upload">
      <input type="file" @change="onFileChange" accept="image/*" />
      <div v-if="imagePreview" class="image-preview">
        <img :src="imagePreview" alt="Image Preview" />
      </div>
      <button v-if="imageFile" @click="uploadImage">Upload Image</button>
    </div>
  </template>
  
<script>
  import {uploadphoto} from '@/services/photos.js';
  export default {
    data() {
      return {
        imageFile: null,
        imagePreview: null,
      };
    },
    methods: {
        onFileChange(event) {
            const file = event.target.files[0];
            if (file && file.type.startsWith("image/")) {
                this.imageFile = file;
                this.imagePreview = URL.createObjectURL(file);
                const reader = new FileReader();
                reader.onload = (e) => {
                // The result is a Base64-encoded string
                this.imageBase64 = e.target.result;
                // If you need to extract just the Base64 part (without the data URL prefix):
                this.imageBase64 = this.imageBase64.split(',')[1];
                };

                reader.readAsDataURL(file);
            } else {
                alert("Please select a valid image file.");
                this.clearImage();
            }
        },


      async uploadImage() {
        console.log("test")
        try{
            let response = await uploadphoto(this.imageBase64)
			console.log('comment added successfully', response);
            this.$router.push({ name: 'profile' });
        }catch(error){
            console.error('Error loading photos:', error);
        }

      },
      clearImage() {
        this.imageFile = null;
        this.imagePreview = null;
      }
    }
  };
  </script>
  
  <style scoped>
  .image-upload {
    margin: 20px;
  }
  
  .image-preview img {
    max-width: 100%;
    max-height: 300px;
    margin-top: 10px;
    border: 1px solid #ddd;
  }
  
  button {
    margin-top: 10px;
    padding: 10px 20px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #0056b3;
  }
  </style>
  