import axios from "axios";

class EndpointsService {
  constructor() {
    // Check vue.config.js for proxy configuration
    this._baseURL = `/api/`;
  }

  async getBuyers(page = 1, limit = 15) {
    try {
      const URL = `${this._baseURL}customers?page=${page}&limit=${limit}`;
      const { data, status } = await axios.get(URL);

      return { data: data.buyers, status, success: true };
    } catch (err) {
      return {
        data: err.response.data,
        status: err.response.status,
        success: false,
      };
    }
  }
}

export default new EndpointsService();
