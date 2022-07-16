/*
"backdrop_sizes": [
      "w45",
      "w92",
      "w154",
      "w185",
      "w300",
      "w500",
      "w780",
      "w1280",
      "w1920",
      "original"
    ],
    "logo_sizes": [
      "w45",
      "w92",
      "w154",
      "w185",
      "w300",
      "w500",
      "original"
    ],
    "poster_sizes": [
      "w45",
      "w92",
      "w154",
      "w185",
      "w300",
      "w342",
      "w500",
      "w780",
      "original"
    ],
    "profile_sizes": [
      "w45",
      "w92",
      "w154",
      "w185",
      "h632",
      "original"
    ],
    "still_sizes": [
      "w45",
      "w92",
      "w154",
      "w185",
      "w300",
      "original"
    ]
 */

async function movieInfo(movie) {
    try {
        const res = await fetch(
            `https://api.themoviedb.org/3/find/${movie}?api_key=0bdcbcb13392d6bd0dea6684fc5be9e8&external_source=imdb_id`
        )
        const result = await res.json()

        if (result["movie_results"].length) {
            return result["movie_results"][0]
        } else {
            alert("Error: movie not found")
        }
    } catch (e) {
        console.log(e)
    }
}

export default {
    name: "Movie info",

    setup() {
        return {title: "movie info"}
    },

    data() {
        return {
            loading: true,
            movie: {},
            additionalInfo: {}
        }
    },

    methods: {
        async fetchMovieInfo() {
            try {
                this.loading = true
                const urlSearchParams = new URLSearchParams(window.location.search);
                const urlParams = Object.fromEntries(urlSearchParams.entries());
                if (!"movieId" in urlParams) {
                    alert("Error: Please specify movie id")
                    return
                }
                if (!"id" in urlParams) {
                    alert("Error: Please specify movie id")
                    return
                }

                this.movie = await movieInfo(urlParams["movieId"])

                const res = await fetch(
                    `/api/movie/${urlParams["id"]}/additionalInfo`
                )
                this.additionalInfo = await res.json()
            } catch (e) {
                console.log(e)
            } finally {
                this.loading = false
            }
        },
        getImageUrlWithWidth(width, path) {
            return `https://image.tmdb.org/t/p/w${width}${path}`
        },
        getImageUrl(width, path) {
            return `https://image.tmdb.org/t/p/original${path}`
        },
        getBackdropImageStyle(url) {
            return `
              background: url("${url}");
              background-size: cover;
              background-position: center;
              height: 400px;
            `
        }
    },

    mounted() {
        this.fetchMovieInfo()
    },

    template: `
    <div>
        <div class="d-flex justify-content-center m-4" v-if="loading">
            <div class="spinner-border" role="status">
                <span class="visually-hidden">Loading...</span>
            </div>
        </div>
        
        <div v-if="!loading">
            <div class="w-100" :style="getBackdropImageStyle(getImageUrlWithWidth(1280, movie.backdrop_path))">
            </div>

            <div class="d-flex justify-content-center mt-4">
                <div class="card mb-3 text-bg-light" style="max-width: 1000px">
                    <div class="row g-0">
                        <div class="col-md-3">
                            <img :src="getImageUrlWithWidth(200, movie.poster_path)" class="img-fluid rounded-start" alt="...">
                        </div>
                        <div class="col-md-8">
                            <div class="card-body">
                            <h5 class="card-title">{{movie.title}}</h5>
                            <p class="card-text">{{movie.overview}}</p>
                            <p class="card-text" style="margin-bottom: 5px !important;">
                                    Rating: {{movie.vote_average}}/10 ({{movie.vote_count}})
                            </p>
                            <div class="progress" style="width: 200px">
                                <div class="progress-bar bg-success" role="progressbar" :style="{width: Math.round(movie.vote_average*10)+'%'}" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100">
                                </div>
                            </div>
                            <p class="card-text mt-4"><small class="text-muted">Released on {{movie.release_date}}</small></p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            
            
            <div class="d-flex justify-content-center mt-4">
                <div class="row w-100" style="max-width: 1000px;">
                
                    <div class="col-6 w-50">
                        <div class="card" style="">
                            <div class="card-header">
                            Genres
                            </div>
                            <ul class="list-group list-group-flush">
                                <li class="list-group-item" v-for="genre of additionalInfo.genres">
                                    {{genre.name}}
                                </li>
                            </ul>
                        </div>
                    </div>
                    
                    <div class="col-6 w-50">
                        <div class="card" style="">
                            <div class="card-header">
                            Productions
                            </div>
                            <ul class="list-group list-group-flush">
                                <li class="list-group-item" v-for="production of additionalInfo.productions">
                                    {{production.name}}
                                </li>
                            </ul>
                        </div>
                    </div>
                
                </div>
            </div> 
    </div>
    `
}