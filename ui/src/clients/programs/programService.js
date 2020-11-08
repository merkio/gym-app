import http from "../http-common";

class ProgramService {
    getAll() {
        return http.get("/programs");
    }

    get(id) {
        return http.get(`/programs/${id}`);
    }

    create(data) {
        return http.post("/programs", data);
    }

    update(id, data) {
        return http.put(`/programs/${id}`, data);
    }

    delete(id) {
        return http.delete(`/programs/${id}`);
    }

    deleteAll() {
        return http.delete(`/programs`);
    }

    findByTitle(title) {
        return http.get(`/programs?title=${title}`);
    }
}

export default new ProgramService();
