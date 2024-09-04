import { getuserid } from './auth-store';
import api from './axios';


export async function getcomments(photoid){
    try{
        let response = await api.get(`/photos/comments/${photoid}`);
        console.log('Response:', response);
        return response.data;
    }catch (e){
        throw new Error(e.toString());
    }
};


export async function addcomment(user,photoid,content){
    try{
        let response = await api.post(`/photos/${photoid}/comments`, {
            Comment:content,
            },{headers :{
                "Authorization" : user
            }});
        console.log('Response:', response);
        return response.data
    }catch (e){
        throw new Error(e.toString())
    }
};

export async function removecomment(user,commentid){
    try{
        let response=await api.delete(`/photos/${commentid}/comments`,{
            headers :{
                "Authorization" : user
            }});
        console.log('Response:', response)
        return response
    }catch(e){
        throw new Error(e)
    }
}

