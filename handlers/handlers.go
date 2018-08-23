package handlers

import (
  "database/sql"
  "net/http"
  "strconv"
	"github.com/gin-gonic/gin"
  "../models"
)

func GetPolls(c *gin.Context) {
  c.JSON(http.StatusOK, models.GetPolls(c.MustGet("db").(*sql.DB)))
}

func UpdatePoll(c *gin.Context) {
   var poll models.Poll

   c.Bind(&poll)

   index, _ := strconv.Atoi(c.Param("index"))

   id, err := models.UpdatePoll(c.MustGet("db").(*sql.DB), index, poll.Name, poll.Upvotes, poll.Downvotes)

   if err == nil {
       c.JSON(http.StatusCreated, gin.H{
           "affected": id,
       })
   }

}
