import { getuserid } from './auth-store';
import api from './axios';


export async function uploadphoto(photocontent){
    let id = await getuserid();
    try{
        let response = await api.post(`/photos`, {
            photo:photocontent,
            },{headers :{
                "Authorization" : id
            }});
        console.log('Response:', response);
        return response.data
    }catch (e){
        console.log(e)
        throw new Error(e.toString())
    }
};
