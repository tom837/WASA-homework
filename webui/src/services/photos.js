import { getuserid } from './auth-store';
import axios from './axios';


export async function uploadphoto(photocontent){
    let id = await getuserid();
    console.log("uploading")
    try{
        console.log("still uploading")
        let response = await axios.post(`/photos`, {
            photo:photocontent,
            },{headers :{
                "Authorization" : id
            }});
        console.log("uploaded")
        console.log('Response:', response);
        return response.data
    }catch (e){
        console.log(e)
        throw new Error(e.toString())
    }
};