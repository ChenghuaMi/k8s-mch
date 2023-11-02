/**
 * @author mch
 */

package response

type PodListItem struct {
	Name string  `json:"name"`
	Ready string  `json:"ready"`
	Status string `json:"status"`
	Restarts int32 `json:"restarts"`
	Age  int32 `json:"age"`
	Ip string `json:"ip"`
	Node string `json:"node"`
}