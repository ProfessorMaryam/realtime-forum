export default async function HomeView() {
    return `
        <div class="block3" id="11">
                <div class="frame">
                    <div class="username">
                        <h3>@m7amd777 <span class="post-date"> • 2025</span></h3>
                    </div>

                    <a href="/thread" style="text-decoration:none; color:black;">
                        <h3>New GTA 6 Leks</h3>
                    </a>
                    <div class="content">

                        <p>
                            bro im just clckbaiting yall
                        </p>
                    </div>
                    <div class="post-meta">
                        <div class="category-flairs">
                            <span>General</span>
                            <span>Gaming</span>
                            <span>Music</span>
                        </div>

                        <div class="post-engagement">
                            <div class="post-like-dislike">
                                <form method="POST" action="/like#{{$i}}">
                                    <input type="hidden" name="direction" value="{{$.CurrentPath}}">
                                    <input type="hidden" name="post_id" value="{{.ID}}">
                                    <button type="submit" name="type" value="like" style="display: inline;"
                                        class="like-btn">
                                        <span class="material-icons">thumb_up</span>
                                        <span>66</span>
                                    </button>
                                    <button type="submit" name="type" value="dislike" style="display: inline;"
                                        class="dislike-btn">
                                        <span class="material-icons">thumb_down</span>
                                        <span>2</span>
                                    </button>
                                </form>
                            </div>


                            <div class="comments">
                                <a class="comment-btn" href="/post/{{.ID}}#comment-input" style="text-decoration:none;">
                                    <span class="material-icons">comment</span>
                                    <span> 23 </span>
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
    `
}