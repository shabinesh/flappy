package main

import "github.com/veandco/go-sdl2/sdl"

/*        (0,0 )        x     -2024->           (w, 0)                    */
/*        *****************************************      y                */
/*        *          *            *               *     |                 */
/*        *          *            *               *     |                 */
/*        *          *   1024x740 *               *     v                 */
/*        *          *            *               *     740               */
/*        *          *            *               *                       */
/*        *****************************************                       */
/*        (0, h)                                  (w, h)                  */

// Camera ...
type Camera struct {
	tex             *sdl.Texture
	cX              int32
	backgroundWidth int32
	speed           int32
}

func (c *Camera) draw() {
	// TODO Scroll background: perception of movement
}
