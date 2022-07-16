export default {
    name: 'Home',

    setup() {
        const title = 'Home Page'
        return {title}
    },

    data() {
        return {
            loading: true,
            searchBy: "title",
            searchText: "",
            response: {
                page: 1,
                pageSize: 20
            },
        }
    },

    mounted() {
        this.fetchMovies()
    },

    methods: {
        fetchMovies() {
            this.loading = true

            const urlSearchParams = new URLSearchParams(window.location.search);
            const urlParams = Object.fromEntries(urlSearchParams.entries());

            if ("searchBy" in urlParams) {
                this.searchBy = urlParams["searchBy"]
            }
            if ("searchText" in urlParams) {
                this.searchText = urlParams["searchText"]
            }
            if ("pageSize" in urlParams) {
                this.response.pageSize = this.pageSize || urlParams["pageSize"]
            }
            if ("page" in urlParams) {
                this.response.page = this.page || urlParams["page"]
            }

            if (!this.searchBy || !this.searchText) {
                this.loading = false
                return
            }

            const url = `/api/search/movies${this.getQueryUrl()}`

            fetch(
                url,
                {headers: {'Content-Type': 'application/json'}}
            )
                .then(async (res) => {
                    this.response = await res.json()
                })
                .catch(err => {
                    alert(err)
                })
                .finally(() => {
                    this.loading = false
                })
        },

        getQueryUrl(page) {
            if (!page) {
                page = this.response.page
            }
            return `?searchText=${this.searchText}&searchBy=${this.searchBy}` +
                `&page=${page}&pageSize=${this.response.pageSize}`
        },

        searchSubmit() {
            this.response.page = 1
            window.location.replace(window.location.origin + this.getQueryUrl())
        }
    },

    template: `
    <div class="container">
        <form class="row align-content-center mb-4" @submit.prevent>
            <div class="col-sm-3 p-0">
                <select class="form-select" aria-label="Default select example" style="height:40px;" v-model="searchBy">
                    <option selected value="title">Movie</option>
                    <option value="genre">Genre</option>
                    <option value="production">Production</option>
                </select>
            </div>
            <div class="mb-3 col-6">
                <input class="form-control" v-model="searchText" placeholder="Search">
            </div>
            <button class="btn btn-primary col-3" style="height: 38px;" @click="searchSubmit">Search</button>
        </form> 
        
        <div class="d-flex justify-content-center m-4" v-if="loading">
            <div class="spinner-border" role="status">
                <span class="visually-hidden">Loading...</span>
            </div>
        </div>
    
        <div v-if="!loading && this.response && this.response.movies">
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col">Title</th>
                        <th scope="col">Release Date</th>
                        <th scope="col">Runtime (m)</th>
                        <th scope="col">Rating Count</th>
                        <th scope="col">Rating</th>
                        <th scope="col">Status</th>
                        <th scope="col">Revenue</th>
                        <th scope="col">Populatiry</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="movie in response.movies">
                        <td><a :href="'/movie' + '?' + 'movieId=' + movie.imdbId + '&id=' + movie.id">{{movie.title}}</a></td>
                        <td>{{new Date(movie.releaseDate).getFullYear()}}</td>
                        <td>{{movie.runtime}}</td>
                        <td>{{parseInt(movie.voteCount).toLocaleString() || "-"}}</td>
                        <td>{{movie.voteAverage || "-"}}</td>
                        <td>{{movie.status || "-"}}</td>
                        <td>{{parseInt(movie.revenue).toLocaleString() || "-"}}</td>
                        <td>{{Math.round(parseFloat(movie.popularity)) || "-"}}</td>
                    </tr>
                </tbody>
            </table>
            
            <nav aria-label="Page navigation example">
                <ul class="pagination justify-content-center">
                    <li class="page-item" :class="{disabled: response.page === 1}">
                        <a class="page-link" :href="getQueryUrl(response.page - 1)">
                            Previous
                        </a>
                    </li>
                    <li class="page-item disabled">
                        <a class="page-link" href="#">
                            {{response.page + ' / ' + Math.ceil(response.totalResults / response.pageSize)}}
                        </a>
                    </li>
<!--                    <li :class="{active: response.page === i}" class="page-item" v-for="i in Math.ceil(response.totalResults/response.pageSize)" :key="i">-->
<!--                        <a class="page-link" :href="getQueryUrl(i)">{{i}}</a>-->
<!--                    </li>-->
                    <li class="page-item" :class="{disabled: Math.ceil(response.totalResults / response.pageSize) === response.page}">
                        <a class="page-link" :href="getQueryUrl(response.page + 1)">
                            Next
                        </a>
                    </li>
                </ul>
            </nav>
        </div>
    </div>
    `,
};
