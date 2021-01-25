import http from "../http-common";

class UserService {
    getAll() {
        return http.get("/users")
            .then(result => result)
            .catch(err => {
                console.error(err);
                return {}
            });
    }

    get(id) {
        return http.get(`/users/${id}`)
            .then(result => result)
            .catch(err => {
                console.error(err);
                return {}
            });
    }

    create(data) {
        return http.post("/users", data)
            .then(result => result)
            .catch(err => {
                console.error(err);
                return {}
            });
    }

    update(id, data) {
        return http.put(`/users/${id}`, data)
            .then(result => result)
            .catch(err => {
                console.error(err);
                return {}
            });
    }

    delete(id) {
        return http.delete(`/users/${id}`)
            .then(result => result)
            .catch(err => {
                console.error(err);
                return {}
            });
    }

    deleteAll() {
        return http.delete(`/users`)
            .then(result => result)
            .catch(err => {
                console.error(err);
                return {}
            });
    }

    findByName(name) {
        return http.get(`/users?name=${name}`)
            .then(result => result)
            .catch(err => {
                console.error(err);
                return {}
            });
    }
}

export default new UserService();
