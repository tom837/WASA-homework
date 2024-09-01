import { getuserid } from './auth-store';
import axios from './axios';


export async function getcomments(photoid){
    try{
        let id = await getuserid();
        let response = await axios.get(`/photos/comments/p12`);
        console.log('Response:', response);
        return response.data;
    }catch (e){
        throw new Error(e.toString());
    }
};
