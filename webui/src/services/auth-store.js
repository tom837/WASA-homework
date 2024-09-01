const token = 'userid';
const storage = sessionStorage;

export function saveuserid(authToken) {
    storage.setItem(token, authToken);
}

export function deleteuserid(){
	const exist = sessionStorage.getItem(token);
	if (exist) {
		storage.removeItem(token);
	}
}

export function getuserid() {
	const user = sessionStorage.getItem(token);
	if (user) {
		return user
	}
	return null
}
