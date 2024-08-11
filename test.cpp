#include "c_middleware.h"

// gcc -o test _main.c libgolang.a -lpthread
int main() {
    InitMiddleware();
    // GetRawData(34.46, 78.41, 34.47, 78.42, 14);
    // GetRoutePoints_return params = GetRoutePoints(113.5439372, 22.2180642, 113.5425177, 22.2252363);
    // printf("total points count: %d\n", params.r0);
    // for (int i = 0; i < params.r0; i++) {
    //     printf("lon-lat: (%lf, %lf)\n", params.r1[i], params.r2[i]);
    // }
    // FreePositionsPointer(params.r1, params.r2);
    // UpdateObstacle(113.416793, 22.158472, "test obstacle");
    UpdateCrater(78.45, 34.49, 5.33, 4.32);
    // double altitude = GetAltitude(78.45, 34.39);
    // printf("altitude at lon-lat: (%lf, %lf) is %lf\n", 78.45, 34.39, altitude);
    CloseMiddleware();
}