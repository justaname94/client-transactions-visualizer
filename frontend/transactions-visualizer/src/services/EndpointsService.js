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

      return { data: data.buyers, success: true, status };
    } catch ({ response: { data, status } }) {
      return {
        data: data.message,
        success: false,
        status
      };
    }
  }

  async loadDate(date) {
    try {
      const URL = `${this._baseURL}load/${date}`;
      const { data, status } = await axios.get(URL);

      return { data: data.success, success: true, status };
    } catch ({ response: { data, status } }) {
      return {
        data: status === 500 ? "An unexpected error ocurred" : data.message,
        success: false,
        status
      };
    }
  }
}

export default new EndpointsService();
